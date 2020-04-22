import { Component, OnInit, NgModule } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Router } from '@angular/router';
import { CommonService } from '../common.service';
import { ToastrService } from 'ngx-toastr';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent implements OnInit {
  registerInformation = {
    registerEmail: '',
    registerPassword: '',
    registerRepeatPassword: '',
    vrc: '',
  };

  emailForm = {
    email: '',
  };

  registerForm = {
    email: '',
    password: '',
    vrc: ''
  };

  constructor(
    private http: HttpClient,
    private router: Router,
    private common: CommonService,
    private toast: ToastrService
  ) { }

  ngOnInit() {
  }

  // 返回注册校验码
  registerFormCheck() {
    if (!this.common.checkEmail(this.registerInformation.registerEmail)) {
      return this.common.emailFormatIncorectFlag;
    }
    if (!this.common.checkPassword(this.registerInformation.registerPassword)) {
      return this.common.passwordFormatIncorectFlag;
    }
    if (!this.common.checkPassword(this.registerInformation.registerRepeatPassword)) {
      return this.common.repeatPasswordFormatIncorectFlag;
    }
    if (this.registerInformation.registerPassword !== this.registerInformation.registerRepeatPassword) {
      return this.common.passwordNotConsistFlag;
    }
    return this.common.corectFlag;
  }
  // 返回校验码
  emailFormCheck() {
    if (!this.common.checkEmail(this.registerInformation.registerEmail)) {
      return this.common.emailFormatIncorectFlag;
    }
    return this.common.corectFlag;
  }

  // 发送验证码
  sendVrc() {
    // 前置校验
    if (this.emailFormCheck() !== this.common.corectFlag) {
      this.toast.error(this.common.parseFlag(this.emailFormCheck()), '提示', {timeOut: 4000});
      return;
    }

    // 数据结构构建
    const header = new HttpHeaders({
      'Content-Type': 'application/x-www-form-urlencoded;charset=UTF-8'
    });
    this.emailForm = {
      email: this.registerInformation.registerEmail,
    };
    const requestHead = { headers: header };
    this.common.reqProto = {
      data: this.emailForm,
      orderBy: '',
      filter: '',
      page: 0,
      pageSize: 0,
    };

    // 发送邮箱验证码请求
    this.http.post(this.common.sendVrcUrl, this.common.reqProto, requestHead).subscribe((res: any) => {
      this.common.replyProto = res;
      if (res.status !== this.common.sendVrcSuccessFlag) {
        this.toast.warning(this.common.replyProto.msg, '提示');
        return;
      }
      this.toast.success(this.common.replyProto.msg, '提示', {timeOut: 4000});
    });
  }


  // 发送Http注册请求
  askForRegister() {
    // 前置校验
    if (this.registerFormCheck() !== this.common.corectFlag) {
      this.toast.error(this.common.parseFlag(this.registerFormCheck()), '提示', {timeOut: 4000});
      return;
    }

    // 数据结构构建
    const header = new HttpHeaders({
      'Content-Type': 'application/x-www-form-urlencoded;charset=UTF-8'
    });
    const requestHead = { headers: header };
    this.registerForm = {
      email: this.registerInformation.registerEmail,
      password: this.registerInformation.registerPassword,
      vrc: this.registerInformation.vrc,
    };
    this.common.reqProto = {
      data: this.registerForm,
      orderBy: '',
      filter: '',
      page: 0,
      pageSize: 0,
    };


    // 发送注册请求
    this.http.post(this.common.registerUrl, this.common.reqProto, requestHead).subscribe((res: any) => {
      this.common.replyProto = res;
      if (this.common.replyProto.status !== this.common.registerSuccessFlag) {
        this.toast.warning(this.common.replyProto.msg, '提示');
        return;
      }

      // 执行操作
      this.common.storeUserPersonalInformation(null);
      this.common.storeUserAccountInformation(null);
      this.router.navigate(['registerResponse']);
      this.toast.success(this.common.replyProto.msg, '提示', {timeOut: 4000});
    });
  }
}
