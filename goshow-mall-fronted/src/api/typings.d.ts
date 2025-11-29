declare namespace API {
  type AddAdminDto = {
    mobile?: string
    nickname?: string
    sex?: number
    username?: string
  }

  type postAdminStatusIdStatusParams = {
    /** 用户id */
    id: number
    /** 用户状态 */
    status: number
  }

  type ResultVoSlideCaptchaVo = {
    code?: number
    data?: SlideCaptchaVo
    msg?: string
  }

  type SlideCaptchaVo = {
    ImageBase64?: string
    TitleHeight?: number
    TitleImageBase64?: string
    TitleWidth?: number
    TitleX?: number
    TitleY?: number
    key?: string
  }

  type UpdateAdminDto = {
    id?: number
    mobile?: string
    nickname?: string
    sex?: number
    username?: string
  }
}
