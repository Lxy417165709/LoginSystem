import { Injectable, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ToastrService } from 'ngx-toastr';
import { HttpClient, HttpHeaders } from '@angular/common/http';
@Injectable({
  providedIn: 'root'
})
export class CommonService {
  // 后端响应数据通信协议
  replyProto = {
    status: 0,
    msg: '',
    data: {},
  };

  // 前端请求数据通信协议
  reqProto = {
    data: {},   // 请求数据
    orderBy: '',  // 排序要求
    filter: '',   // 筛选条件
    page: 0,      // 分页
    pageSize: 0,  // 分页大小
  };

  // 用户账户信息结构
  userAccountInformation = {
    userId: -1,
    userEmail: '',
    userPassword: '',
  };

  // 用户个人信息结构
  userPersonalInformation = {
    userId: -1,
    photoData: '',
    userPhoto: '',
    userName: '',
    userSex: '',
    userContactPhone: '',
    userContactEmail: '',
    userBirthday: 0,
  };

  // localstore 存储键映射
  keyOfUai = 'userAccountInformation';
  keyOfUpi = 'userPersonalInformation';


  // 图标、图片路径
  incorrectIcoPath = 'assets/img/incorrect.ico';
  correctIcoPath = 'assets/img/correct.ico';
  defaultImgPath = 'assets/img/default.jpg';  // 默认图片路径
  closeIcoUrl = '/assets/img/close.ico';

  // 校验正则表达式
  EmailReg = /^[a-zA-Z0-9_.-]+@[a-zA-Z0-9-]+(\.[a-zA-Z0-9-]+)*\.[a-zA-Z0-9]{2,6}$/;
  phoneReg = /^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\d{8}$/;
  passWordReg = /^[a-zA-Z0-9_.]{8,15}$/;
  vrcReg = /^[0-9]{6}$/;

  // 标志
  loginSuccessFlag = 1;
  registerSuccessFlag = 1;
  sendVrcSuccessFlag = 1;
  updatePhotoSuccessFlag = 1;
  getUpiSuccessFlag = 1;
  getUaiSuccessFlag = 1;
  updateUpiSuccessFlag = 1;

  // 校验标志
  corectFlag = 1;
  emailFormatIncorectFlag = -1;
  passwordFormatIncorectFlag = -2;
  repeatPasswordFormatIncorectFlag = -3;
  passwordNotConsistFlag = -4;
  vrcFormatIncorectFlag = -5;
  phoneIncorectFlag = -6;
  notPhotoFlag = -7;
  photoSizeTooBigFlag = -8;
  photoEmptyFlag = -9;
  notNewPhotoFlag = -10;
  birthdayIncorectFlag = -11;
  sexIncorectFlag = -12;
  userNameIncorectFlag = -13;
  // 图片
  photoMaxSize = 0.5 * 1024 * 1024; // 0.5M

  // 接口
  loginUrl = '/server/login';
  registerUrl = '/server/register';
  sendVrcUrl = '/server/registerVrc/send';
  updatePhotoUrl = '/server/updatePhoto';
  getUaiUrl = '/server/getUai';
  getUpiUrl = '/server/getUpi';
  updateUpiUrl = '/server/updateUserPersonalInformation';

  // 男女标志
  manFlag = 1;
  womanFlag = 2;

  // 提示位置选项
  /*
  toast-top-left  顶端左边
  toast-top-right    顶端右边
  toast-top-center  顶端中间
  toast-top-full-width 顶端，宽度铺满整个屏幕
  toast-bottom-right
  toast-bottom-left
  toast-bottom-center
  toast-bottom-full-width
  */



  // 解析校验标志
  // 解析校验码
  parseFlag(flag: number) {
    // let tipStr = '';
    // const resultFlag = this.registerFormIsOK();
    if (flag === this.corectFlag) {
      return '';
    }
    if (flag === this.phoneIncorectFlag) {
      return '手机格式有误!';
    }
    if (flag === this.emailFormatIncorectFlag) {
      return '邮箱格式有误!';
    }
    if (flag === this.passwordFormatIncorectFlag) {
      return '密码格式有误!';
    }
    if (flag === this.repeatPasswordFormatIncorectFlag) {
      return '重复密码的格式有误!';
    }
    if (flag === this.passwordNotConsistFlag) {
      return '两次输入的密码不一致';
    }
    if (flag === this.vrcFormatIncorectFlag) {
      return '验证码格式有误';
    }
    if (flag === this.notPhotoFlag) {
      return '该文件不是图片';
    }
    if (flag === this.photoSizeTooBigFlag) {
      return '你选中的图片太大了,图片最大只能为 500KB(0.5MB)';
    }
    if (flag === this.photoEmptyFlag) {
      return '你还没有选择任何图片';
    }
    if (flag === this.notNewPhotoFlag) {
      return '你还没有选择任何新图片';
    }
    if (flag === this.userNameIncorectFlag) {
      return '用户名格式有误!';
    }
    if (flag === this.sexIncorectFlag) {
      return '性别信息错误';
    }
    if (flag === this.birthdayIncorectFlag) {
      return '您选择的生日信息有误!';
    }
    return '';
  }


  // 获取用户头像url
  getUserPhotoUrl() {
    return 'data:image/jpg;base64,' + this.userPersonalInformation.photoData;
  }

  // 这个是抄别人的，把dataUrl转换为Blob
  dataURLtoBlob(dataurl) {
    const arr = dataurl.split(',');
    const mime = arr[0].match(/:(.*?);/)[1];
    const bstr = atob(arr[1]);
    let n = bstr.length;
    const u8arr = new Uint8Array(n);
    while (n--) {
        u8arr[n] = bstr.charCodeAt(n);
    }
    return new Blob([u8arr], { type: mime });
  }

  // 获取文件后缀名 (不包括小数点)
  getFileSuffix(file) {
    const index = file.name.lastIndexOf('.');
    return file.name.substr(index + 1);
  }

  // 判断文件是否是图片文件
  isPhotoFile(file) {
    const suffix = this.getFileSuffix(file).toLowerCase();
    return suffix === 'gif' || suffix === 'bmp' || suffix === 'jpg' || suffix === 'jpeg' || suffix === 'png';
  }

  // 检查用户名格式是否合法
  checkUsername(username: string): boolean {
    return this.checkEmail(username);
  }

  // 检查验证码格式是否合法
  checkVrc(vrc: string): boolean {
    return this.vrcReg.test(vrc);
  }

  // 检查登录密码格式是否合法
  checkPassword(password: string) {
    return this.passWordReg.test(password);
  }

  // 检查性别是否合法
  checkPersonalSex(sex) {
    return sex === this.manFlag || sex === this.womanFlag;
  }
  // 检查注册手机格式是否满足要求
  checkPhone(phone: string) {
    return this.phoneReg.test(phone);
  }

  // 检查注册邮箱格式是否满足要求
  checkEmail(email: string): boolean {
    return this.EmailReg.test(email) && email.length <= 30;
  }


  // 检查个人用户名格式是否满足要求
  checkPersonalName(personalName: string) {
    return personalName.length >= 1 && personalName.length <= 10;
  }

  // 检查个人生日是否满足要求
  checkPersonalBirthday(personalBirthday) {
    return !isNaN(personalBirthday) && personalBirthday < new Date().getTime();
  }

  // 通过布尔值返回正确or错误的Ico的URL
  getIcoUrl(flag: boolean): string {
    if (flag === false) {
      return this.incorrectIcoPath;
    } else {
      return this.correctIcoPath;
    }
  }

  // 检测登录状态 (false表示没登陆,true表示已登录)
  loginStateDetection(): boolean {
    // 检测浏览器是否支持storage
    if (!sessionStorage || !localStorage) {
      console.warn('本浏览器不支持storage!\n');
      return false;
    }
    // 从session storage中解析获取用户账户信息
    return this.getUserAccountInformation() !== null;
  }

  // 从session storage中解析中用户账户信息
  getUserAccountInformation() {
    return JSON.parse(sessionStorage.getItem(this.keyOfUai));
  }

  // 从session storage中解析中用户个人信息 (包括了头像的安全链接获取)
  getUserPersonalInformation() {
    return JSON.parse(sessionStorage.getItem(this.keyOfUpi));
  }

  getImgSrcPhoto() {
    if (this.userPersonalInformation.photoData === '') {
      return this.defaultImgPath;
    }
    return 'data:image/jpg;base64,' + this.userPersonalInformation.photoData;
  }

  // 获得纯净的base64编码(不要 'data:image/jpg;base64,')
  getPureBase64(notPureBase64: string) {
    return notPureBase64.substr(notPureBase64.indexOf(',') + 1);
  }

  // 将用户账户信息存储到session storage中
  storeUserAccountInformation(userAccountInformation) {
    sessionStorage.setItem(this.keyOfUai, JSON.stringify(userAccountInformation));
  }
  storeUserPersonalInformation(userPersonalInformation) {
    sessionStorage.setItem(this.keyOfUpi, JSON.stringify(userPersonalInformation));
  }

  // 删除用户信息 (包括用户个人信息和用户账户信息)
  clearUserInformation() {
    sessionStorage.clear();
  }

  // 将时间戳转换为对应的时间字符串 (单位 unix *1e3)
  timestampToTimeString(transTime: number) {
    const date = new Date(transTime);
    const year = date.getFullYear();
    const month = date.getMonth() + 1;
    const d = date.getDate();
    let dString = '' + d;
    let mString = '' + month;
    if (month < 10) {
      mString = '0' + mString;
    }
    if (d < 10) {
      dString = '0' + dString;
    }

    return year + '-' + mString + '-' + dString ;
  }


  getUpi() {
    // 数据结构构建
    const header = new HttpHeaders({
      'Content-Type': 'application/x-www-form-urlencoded;charset=UTF-8'
    });
    const requestHead = { headers: header };

    // 通过token 获取用户信息
    this.http.post(this.getUpiUrl, this.reqProto, requestHead).subscribe((res: any) => {
      console.log(res);
      this.replyProto = res;
      if (res.status !== this.getUpiSuccessFlag) {
        this.toast.warning(this.replyProto.msg, '提示');
        this.router.navigate(['login']);
        return;
      }

      // 存储用户账户信息
      this.userPersonalInformation = res.data;
      this.storeUserPersonalInformation(this.userPersonalInformation);
      this.toast.success(this.replyProto.msg, '提示', {positionClass: 'toast-bottom-right'});
    });
  }

  getUai() {
    // 数据结构构建
    const header = new HttpHeaders({
      'Content-Type': 'application/x-www-form-urlencoded;charset=UTF-8'
    });
    const requestHead = { headers: header };

    // 通过token 获取用户信息
    this.http.post(this.getUaiUrl, this.reqProto, requestHead).subscribe((res: any) => {
      console.log(res);
      this.replyProto = res;
      if (res.status !== this.getUaiSuccessFlag) {
        this.toast.warning(this.replyProto.msg, '提示');
        this.router.navigate(['login']);
        return;
      }

      // 存储用户账户信息
      this.userAccountInformation = res.data;
      this.storeUserAccountInformation(this.userAccountInformation);
      this.toast.success(this.replyProto.msg, '提示', {positionClass: 'toast-bottom-right'});
    });
  }





  constructor(
    private http: HttpClient,
    private router: Router,
    private toast: ToastrService
  ) { }

  OnInit() {
  }


}
