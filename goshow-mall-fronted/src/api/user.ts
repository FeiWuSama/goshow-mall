// @ts-ignore
/* eslint-disable */
import request from '@/utils/request'

/** 获取滑块验证码 GET /api/user/captcha/slide */
export async function getUserCaptchaSlide(options?: { [key: string]: any }) {
  return request<API.ResultVoSlideCaptchaVo>('/api/user/captcha/slide', {
    method: 'GET',
    ...(options || {}),
  })
}
