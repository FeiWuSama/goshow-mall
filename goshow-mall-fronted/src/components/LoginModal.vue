<script setup lang="ts">
import { ref, onMounted, watch, nextTick } from 'vue'
import { Form, Input, Button, Tabs, Space, message, Modal } from 'ant-design-vue'
import { UserOutlined, LockOutlined, PhoneOutlined, QrcodeOutlined } from '@ant-design/icons-vue'
import { getUserCaptchaSlide } from '@/api/user.ts'

interface Props {
  visible: boolean
}

const props = defineProps<Props>()

const emit = defineEmits<{
  'update:visible': [value: boolean]
  'login-success': [token: string]
}>()

// 账号密码登录
const accountForm = ref({
  username: '',
  password: '',
})

// 验证码登录
const codeForm = ref({
  phone: '',
  code: '',
})

const isLoading = ref(false)
const activeTab = ref('account')
const countDown = ref(0)

// 验证码相关状态
const captchaModal = ref(false)
const captchaData = ref<any>(null)
const captchaKey = ref('')
const captchaSliderPosition = ref(0)
const isDragging = ref(false)
const captchaVerifying = ref(false)
const sliderContainer = ref<HTMLElement | null>(null)
const captchaImageRef = ref<HTMLImageElement | null>(null)
const yScaleRatio = ref(1) // Y轴缩放比例

// 关闭弹窗
const handleClose = () => {
  emit('update:visible', false)
  resetForms()
}

// 重置表单
const resetForms = () => {
  accountForm.value = { username: '', password: '' }
  codeForm.value = { phone: '', code: '' }
  activeTab.value = 'account'
  countDown.value = 0
}

// 账号密码登录
const handleAccountLogin = async () => {
  if (!accountForm.value.username || !accountForm.value.password) {
    message.error('请填写用户名和密码')
    return
  }

  isLoading.value = true
  // try {
  //   const response = await loginByAccount({
  //     username: accountForm.value.username,
  //     password: accountForm.value.password,
  //   })
  //   message.success('登录成功')
  //   console.log('登录响应:', response)
  //   // TODO: 保存token，触发登录成功事件
  //   emit('login-success', 'token')
  //   handleClose()
  // } catch (error) {
  //   message.error('登录失败')
  // } finally {
  //   isLoading.value = false
  // }
}

// 获取验证码
const handleGetCode = async () => {
  if (!codeForm.value.phone) {
    message.error('请输入手机号')
    return
  }

  try {
    // 先获取滑块验证码
    let result = await getUserCaptchaSlide()
    if (result.code === 20000 && result.data) {
      captchaData.value = result.data
      captchaKey.value = result.data.key
      captchaSliderPosition.value = 0 // 初始位置设置为0
      captchaModal.value = true // 打开滑块验证码弹窗
    } else {
      message.error('获取验证码失败')
    }
  } catch (error) {
    message.error('获取验证码失败')
  }
}

// 验证码登录
const handleCodeLogin = async () => {
  if (!codeForm.value.phone || !codeForm.value.code) {
    message.error('请填写手机号和验证码')
    return
  }

  isLoading.value = true
  // try {
  //   const response = await loginByCode({
  //     phone: codeForm.value.phone,
  //     code: codeForm.value.code,
  //   })
  //   message.success('登录成功')
  //   console.log('登录响应:', response)
  //   emit('login-success', 'token')
  //   handleClose()
  // } catch (error) {
  //   message.error('登录失败')
  // } finally {
  //   isLoading.value = false
  // }
}

// 验证滑块
const handleVerifyCaptcha = async () => {
  captchaVerifying.value = true
  // try {
  //   // 计算拼图的实际位置，即滑块位置加上拼图初始位置
  //   const puzzleX = captchaSliderPosition.value + captchaData.value.TitleWidth
  //
  //   const response = await verifyVerificationCode({
  //     key: captchaKey.value,
  //     x: puzzleX,
  //     y: captchaData.value.TitleY, // 使用接口返回的Y坐标
  //   })
  //   if (response.code === 20000) {
  //     // message.success('验证成功，正为你发送短信验证码...')
  //     // captchaModal.value = false
  //     // // 验证成功后，获取短信验证码
  //     // await sendVerificationCode({
  //     //   phone: codeForm.value.phone,
  //     // })
  //     // message.success('短信验证码已发送')
  //     // startCountDown()
  //   } else {
  //     message.error('验证失败，请重试')
  //     // 重置滑块位置
  //     captchaSliderPosition.value = 0
  //   }
  // } catch (error) {
  //   message.error('验证失败')
  //   // 重置滑块位置
  //   captchaSliderPosition.value = 0
  // } finally {
  //   captchaVerifying.value = false
  // }
}

// 自动校验滑块位置
const checkSliderPosition = () => {
  if (!captchaData.value) return

  // 计算拼图的实际位置，即滑块位置加上拼图初始位置
  const puzzleX = captchaSliderPosition.value + captchaData.value.TitleWidth
  // 目标位置是缺口位置
  const targetX = captchaData.value.TitleX
  const tolerance = 5 // 允许的误差范围

  // 如果拼图位置与目标位置接近，返回true表示可以验证
  return Math.abs(puzzleX - targetX) <= tolerance
}

// 关闭slider验证码弹窗
const handleCloseCaptchaModal = () => {
  captchaModal.value = false
  captchaSliderPosition.value = 0
}

// 处理滑块按下事件
const handleSliderMouseDown = (event: MouseEvent) => {
  event.preventDefault()
  isDragging.value = true

  // 在开始拖动时缓存容器信息，避免拖动过程中频繁计算
  cacheContainerInfo()

  // 重置lastPosition
  lastPosition = captchaSliderPosition.value

  // 添加全局鼠标事件监听
  document.addEventListener('mousemove', handleGlobalMouseMove)
  document.addEventListener('mouseup', handleGlobalMouseUp)
}

// 用于requestAnimationFrame的位置更新函数
let animationFrameId: number
let lastPosition: number = -1
let containerInfo: { width: number; left: number } | null = null

// 缓存容器信息
const cacheContainerInfo = () => {
  if (!sliderContainer.value) return
  const rect = sliderContainer.value.getBoundingClientRect()
  containerInfo = {
    width: rect.width,
    left: rect.left
  }
}

// 全局鼠标移动事件处理
const handleGlobalMouseMove = (event: MouseEvent) => {
  if (!isDragging.value) return

  event.preventDefault()

  // 使用requestAnimationFrame优化性能，避免频繁更新
  if (animationFrameId) return

  animationFrameId = requestAnimationFrame(() => {
    // 只在缓存为空时重新获取容器信息
    if (!containerInfo) {
      cacheContainerInfo()
    }

    if (!containerInfo || !captchaData.value) {
      animationFrameId = 0
      return
    }

    // 使用缓存的容器信息，减少getBoundingClientRect调用
    const titleWidth = captchaData.value.TitleWidth * 0.5

    // 计算新位置，确保滑块不会超出容器边界
    let newPosition = event.clientX - containerInfo.left - titleWidth
    newPosition = Math.max(0, Math.min(newPosition, containerInfo.width - 50))

    // 只有位置发生变化超过1px时才更新，避免极小的波动引起渲染
    if (Math.abs(newPosition - lastPosition) >= 1) {
      captchaSliderPosition.value = newPosition
      lastPosition = newPosition
    }

    animationFrameId = 0
  })
}

// 全局鼠标释放事件处理
const handleGlobalMouseUp = () => {
  isDragging.value = false

  // 取消可能存在的动画帧
  if (animationFrameId) {
    cancelAnimationFrame(animationFrameId)
    animationFrameId = 0
  }

  // 清除缓存的容器信息
  containerInfo = null

  // 移除全局鼠标事件监听
  document.removeEventListener('mousemove', handleGlobalMouseMove)
  document.removeEventListener('mouseup', handleGlobalMouseUp)

  // 使用requestAnimationFrame确保最后一次位置更新完成后再验证
  requestAnimationFrame(() => {
    setTimeout(() => {
      handleVerifyCaptcha()
    }, 500)
  })
}

// 飞书扫码登录
const handleFeishuLogin = async () => {
  // try {
  //   const response = await getFeishuLoginUrl()
  //   console.log('飞书登录URL:', response)
  //   message.info('飞书扫码登录准备完毕...')
  // } catch (error) {
  //   message.error('获取飞书登录地址失败')
  // }
}

// 开始倒计时
const startCountDown = () => {
  countDown.value = 60
  const timer = setInterval(() => {
    countDown.value--
    if (countDown.value <= 0) {
      clearInterval(timer)
    }
  }, 1000)
}

// 绘制波浪
let animationId: number
const drawWaves = (ctx: CanvasRenderingContext2D, canvas: HTMLCanvasElement) => {
  let time = 0

  const draw = () => {
    // 获取实际的canvas尺寸
    const width = canvas.width
    const height = canvas.height

    // 清空画布
    ctx.clearRect(0, 0, width, height)

    // 设置渐变
    const gradient = ctx.createLinearGradient(0, 0, 0, height)
    gradient.addColorStop(0, 'rgba(24, 144, 255, 0.8)')
    gradient.addColorStop(1, 'rgba(24, 144, 255, 0.1)')

    // 绘制第一条波浪
    ctx.beginPath()
    ctx.moveTo(0, height * 0.5)
    for (let x = 0; x <= width; x += 5) {
      const y =
        height * 0.5 +
        Math.sin((x / width) * Math.PI * 2 + time * 0.05) * 15 +
        Math.sin((x / width) * Math.PI + time * 0.02) * 10
      ctx.lineTo(x, y)
    }
    ctx.lineTo(width, height)
    ctx.lineTo(0, height)
    ctx.fillStyle = gradient
    ctx.fill()

    // 绘制第二条波浪
    ctx.beginPath()
    ctx.moveTo(0, height * 0.55)
    for (let x = 0; x <= width; x += 5) {
      const y =
        height * 0.55 +
        Math.sin((x / width) * Math.PI * 2 + time * 0.04) * 18 +
        Math.sin((x / width) * Math.PI + time * 0.03) * 12
      ctx.lineTo(x, y)
    }
    ctx.lineTo(width, height)
    ctx.lineTo(0, height)
    ctx.fillStyle = 'rgba(24, 144, 255, 0.3)'
    ctx.fill()

    time++
    animationId = requestAnimationFrame(draw)
  }

  draw()
}

// 初始化Canvas动画
const initWaveAnimation = async () => {
  await nextTick()
  const canvas = document.querySelector('.login-wave-canvas') as HTMLCanvasElement
  if (canvas) {
    // 设置canvas的实际像素大小
    const rect = canvas.getBoundingClientRect()
    canvas.width = rect.width
    canvas.height = rect.height

    const ctx = canvas.getContext('2d')
    if (ctx) {
      drawWaves(ctx, canvas)
    }
  }
}

onMounted(() => {
  initWaveAnimation()
})

watch(
  () => props.visible,
  (newVal) => {
    if (newVal) {
      initWaveAnimation()
    }
  },
)
</script>

<template>
  <Modal
    :open="props.visible"
    :footer="null"
    @cancel="handleClose"
    width="420px"
    centered
    :mask-style="{ backdropFilter: 'blur(5px)' }"
  >
    <!-- 波浪背景 canvas -->
    <canvas class="login-wave-canvas"></canvas>

    <div class="login-modal-content">
      <h1 class="login-title">GoshowMall</h1>
      <p class="login-subtitle">欢迎登录</p>

      <Tabs v-model:activeKey="activeTab" animated class="login-tabs" size="small">
        <Tabs.TabPane key="account" tab="账号登录">
          <Form layout="vertical" @finish="handleAccountLogin">
            <Form.Item label="用户名" required>
              <Input
                v-model:value="accountForm.username"
                placeholder="请输入用户名或邮箱"
                size="large"
                allow-clear
              >
                <template #prefix>
                  <UserOutlined />
                </template>
              </Input>
            </Form.Item>

            <Form.Item label="密码" required>
              <Input.Password
                v-model:value="accountForm.password"
                placeholder="请输入密码"
                size="large"
              >
                <template #prefix>
                  <LockOutlined />
                </template>
              </Input.Password>
            </Form.Item>

            <Form.Item>
              <Button type="primary" html-type="submit" size="large" block :loading="isLoading">
                登录
              </Button>
            </Form.Item>
          </Form>
        </Tabs.TabPane>

        <!-- 验证码登录 -->
        <Tabs.TabPane key="code" tab="验证码登录">
          <Form layout="vertical" @finish="handleCodeLogin">
            <Form.Item label="手机号" required>
              <Input
                v-model:value="codeForm.phone"
                placeholder="请输入手机号"
                size="large"
                allow-clear
              >
                <template #prefix>
                  <PhoneOutlined />
                </template>
              </Input>
            </Form.Item>

            <Form.Item label="验证码" required>
              <Space style="width: 100%; gap: 8px">
                <Input
                  v-model:value="codeForm.code"
                  placeholder="请输入验证码"
                  size="large"
                  style="flex: 1"
                  allow-clear
                />
                <Button
                  type="default"
                  size="large"
                  :disabled="countDown > 0"
                  @click="handleGetCode"
                >
                  {{ countDown > 0 ? `${countDown}s` : '获取验证码' }}
                </Button>
              </Space>
            </Form.Item>

            <Form.Item>
              <Button type="primary" html-type="submit" size="large" block :loading="isLoading">
                登录
              </Button>
            </Form.Item>
          </Form>
        </Tabs.TabPane>

        <!-- 飞书扫码登录 -->
        <Tabs.TabPane key="feishu" tab="飞书扫码">
          <div class="feishu-login-container">
            <div class="qrcode-placeholder">
              <QrcodeOutlined style="font-size: 48px; color: #1890ff" />
            </div>
            <p style="text-align: center; color: #666; margin-top: 16px">使用飞书扫描二维码登录</p>
            <Button
              type="primary"
              size="large"
              block
              @click="handleFeishuLogin"
              style="margin-top: 16px"
            >
              获取登录二维码
            </Button>
          </div>
        </Tabs.TabPane>
      </Tabs>

      <!-- 底部链接 -->
      <div class="login-footer">
        <span>没有账号？</span>
        <a href="#" @click.prevent>立即注册</a>
        <span class="divider">|</span>
        <a href="#" @click.prevent>忘记密码？</a>
      </div>
    </div>
  </Modal>

  <!-- 滑块验证码弹窗 -->
  <Modal
    v-model:open="captchaModal"
    title="人机正常验证"
    :footer="null"
    @close="handleCloseCaptchaModal"
    centered
    width="350px"
  >
    <div class="captcha-container">
      <!-- 滑块验证码图片 -->
      <div class="captcha-image-wrapper" v-if="captchaData">
        <!-- 背景图 -->
        <img
          :src="captchaData.ImageBase64"
          alt="验证码"
          class="captcha-image"
        />
        <!-- 拼图 -->
        <img
          v-if="captchaData.TitleImageBase64"
          :src="captchaData.TitleImageBase64"
          alt="拼图"
          class="captcha-puzzle"
          :style="{
            left: captchaSliderPosition + 'px',
            top: captchaData.TitleY + 'px',
          }"
        />
      </div>
      <!-- 滑块轨道 -->
      <div class="captcha-slider-container" ref="sliderContainer">
        <div
          class="captcha-slider-handle"
          :class="{ dragging: isDragging }"
          :style="{ left: captchaSliderPosition + 'px' }"
          @mousedown="handleSliderMouseDown"
        >
          <span class="slider-icon">→</span>
        </div>
      </div>
    </div>
  </Modal>
</template>

<style scoped>
:deep(.ant-modal-content) {
  position: relative;
  overflow: hidden;
  background: transparent;
}

:deep(.ant-modal-body) {
  position: relative;
  padding: 0;
  overflow: hidden;
  background: #fff;
}

.login-wave-canvas {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 120px;
  z-index: 0;
}

.login-modal-content {
  padding: 20px 24px;
  position: relative;
  z-index: 1;
}

.login-title {
  text-align: center;
  font-size: 28px;
  font-weight: bold;
  color: #1890ff;
  margin: 0 0 8px 0;
}

.login-subtitle {
  text-align: center;
  color: #666;
  margin: 0 0 24px 0;
  font-size: 14px;
}

.login-tabs {
  margin-bottom: 16px;
}

/* 简化Tab样式 */
:deep(.ant-tabs-nav) {
  margin: 16px !important;
  padding: 0 !important;
}

:deep(.ant-tabs-nav-wrap) {
  overflow: visible !important;
}

:deep(.ant-tabs-nav::before) {
  display: none !important;
}

:deep(.ant-tabs-nav-list) {
  width: 100% !important;
  justify-content: space-around;
}

:deep(.ant-tabs-tab) {
  margin: 4px !important;
  font-size: 14px !important;
  font-weight: 500 !important;
  flex: 0 0 auto !important;
}

:deep(.ant-tabs-tab-active .ant-tabs-tab-btn) {
  margin: 4px !important;
  color: #1890ff !important;
  font-weight: 600 !important;
}

:deep(.ant-tabs-ink-bar) {
  height: 3px !important;
  background-color: #1890ff !important;
  border-radius: 2px;
  bottom: -3px !important; /* 使下划线离文字编距壳异的三像姓等的上面空隙 */
  transition: all 0.3s ease !important;
}

/* 控制Tab内容滑动 */
:deep(.ant-tabs-content-holder) {
  /* 使用CSS transitions并限制动画在弹窗范围内 */
  overflow: hidden;
  position: relative;
}

:deep(.ant-tabs-tabpane) {
  animation: none !important;
  /* 禁用Ant Design默认动画 */
}

:deep(.ant-tabs-tab-active) {
  animation: slideIn 0.3s ease-in-out forwards;
}

@keyframes slideIn {
  from {
    opacity: 0.8;
    transform: translateX(10px);
  }
  to {
    opacity: 1;
    transform: translateX(0);
  }
}

.feishu-login-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 40px 20px;
  min-height: 300px;
}

.qrcode-placeholder {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 200px;
  height: 200px;
  border: 2px dashed #d9d9d9;
  border-radius: 8px;
  background-color: #fafafa;
}

.login-footer {
  text-align: center;
  margin-top: 24px;
  font-size: 14px;
  color: #666;
}

.login-footer a {
  color: #1890ff;
  text-decoration: none;
  margin: 0 4px;
  transition: color 0.3s;
}

.login-footer a:hover {
  color: #40a9ff;
}

.divider {
  margin: 0 8px;
  color: #ddd;
}

/* 滑块验证码样式 */
.captcha-container {
  display: flex;
  flex-direction: column;
  width: 100%;
  gap: 16px;
}

.captcha-image-wrapper {
  position: relative;
  overflow: hidden;
  border-radius: 4px;
  /* 添加GPU加速 */
  transform: translateZ(0);
  will-change: transform;
}

.captcha-image {
  position: relative;
  width: 100%;
}

.captcha-puzzle {
  position: absolute;
  z-index: 2;
  transition: none; /* 拖动时禁用过渡效果，提升性能 */
  pointer-events: none;
  width: auto;
  height: auto;
  will-change: left; /* 告诉浏览器我们将改变left属性，优化渲染 */
  /* 添加GPU加速 */
  transform: translateZ(0);
}

.captcha-slider-container {
  position: relative;
  width: 100%;
  height: 50px;
  background-color: #f7f9fa;
  border-radius: 4px;
  margin-bottom: 16px;
  border: 1px solid #e0e0e0;
  box-shadow: inset 0 0 5px rgba(0, 0, 0, 0.05);
  /* 添加GPU加速 */
  transform: translateZ(0);
}

.captcha-slider-handle {
  position: absolute;
  top: 50%;
  left: 0;
  width: 50px;
  height: 50px;
  background-color: #fff;
  border-radius: 4px;
  cursor: grab;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #666;
  font-weight: bold;
  transition: all 0.05s ease-out;
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.15);
  user-select: none;
  z-index: 10;
  transform: translateY(-50%) translateZ(0); /* 合并transform属性并添加GPU加速 */
  border: 1px solid #e0e0e0;
  will-change: left, background-color, box-shadow; /* 优化渲染 */
}

.captcha-slider-handle:hover {
  background-color: #f0f8ff;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
}

.captcha-slider-handle.dragging {
  cursor: grabbing;
  background-color: #e6f7ff;
  box-shadow: 0 6px 12px rgba(0, 0, 0, 0.25);
}

.slider-icon {
  font-size: 25px;
  color: #1890ff;
}
</style>
