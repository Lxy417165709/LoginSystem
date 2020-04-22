import { Component, OnInit, NgModule } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Router } from '@angular/router';
import { CommonService } from '../common.service';
import { ToastrService } from 'ngx-toastr';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  // 登录表单
  loginForm = {
    email: '',
    password: '',
  };

  constructor(
    private http: HttpClient,
    private router: Router,
    private common: CommonService,
    private toast: ToastrService,
  ) { }

  ngOnInit() {
    this.toast.info('欢迎登陆', '小提示');
  }

  // 返回登录表单校验码，0表示正确
  loginFormCheck() {
    if (!this.common.checkEmail(this.loginForm.email)) {
      return this.common.emailFormatIncorectFlag;
    }
    if (!this.common.checkPassword(this.loginForm.password)) {
      return this.common.passwordFormatIncorectFlag;
    }
    return this.common.corectFlag;
  }

  // 发送Http登录请求
  askForLogin() {
    // 检测表单信息
    if (this.loginFormCheck() !== this.common.corectFlag) {
      this.toast.error(this.common.parseFlag(this.loginFormCheck()), '错误提示');
      return ;
    }

    // 数据结构构建
    const header = new HttpHeaders({
      'Content-Type': 'application/x-www-form-urlencoded;charset=UTF-8'
    });
    const requestHead = { headers: header };
    this.common.reqProto = {
      data: this.loginForm,
      orderBy: '',
      filter: '',
      page: 0,
      pageSize: 0,
    };

    // 发送请求
    this.http.post(this.common.loginUrl, this.common.reqProto, requestHead).subscribe((res: any) => {
      this.common.replyProto = res;
      if (this.common.replyProto.status !== this.common.loginSuccessFlag) {
        this.toast.warning(this.common.replyProto.msg, '登录提示');
        return;
      }

      // 清空缓存
      this.common.storeUserPersonalInformation(null);
      this.common.storeUserAccountInformation(null);
      this.router.navigate(['home']);
      this.toast.success(this.common.replyProto.msg, '登录提示', {positionClass: 'toast-bottom-right'});
    });
  }
}
