import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { CommonService } from '../common.service';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { DomSanitizer } from '@angular/platform-browser';
import { ToastrService } from 'ngx-toastr';
@Component({
  selector: 'app-personal-information',
  templateUrl: './personal-information.component.html',
  styleUrls: ['./personal-information.component.css']
})
export class PersonalInformationComponent implements OnInit {
  openState = 0; // 用于显示选择照片子组件   0表示隐藏子组件,1表示打开

  constructor(
    private common: CommonService,
    private http: HttpClient,
    private toast: ToastrService,
  ) { }

  ngOnInit() {
    // 初始化信息
    this.openState = 0;
    const uai = this.common.getUserAccountInformation();
    const upi = this.common.getUserPersonalInformation();
    if (uai !== null && upi !== null) {
      this.common.userAccountInformation = uai;
      this.common.userPersonalInformation = upi;
    } else {
      this.common.getUai();
      this.common.getUpi();
    }
  }

  // 打开选择照片的窗口(该窗口是一个组件)
  openSelectPictureWindow() {
    this.openState = 1;
  }

  // 从子组件中获取消息 (子组件会传送一个 0 的值过来，表示关闭子组件)
  getMsgFromSon(Msg) {
    // 表示关闭子组件
    if (Msg.type === 0) {
      this.openState = Msg.data;
      return ;
    }
    // 表示更新用户个人信息
    if (Msg.type === 1) {
      this.common.userPersonalInformation = Msg.data;  // 已规定好，子组件把更新的目的信息交予父组件
      this.askForChangingUserPersonalInformation();
      return ;
    }
  }


  // 由于用户个人信息的生日保存形式是时间戳，因此用户修改个人生日的时候也要修改相应的时间戳
  listenUserBirthdayChanged(e) {
    this.common.userPersonalInformation.userBirthday = new Date(e.target.value).getTime();
  }

  // 个人信息表单
  personalUserFormIsOK() {
    if (!this.common.checkPersonalName(this.common.userPersonalInformation.userName)) {
      return this.common.userNameIncorectFlag;
    }
    if (!this.common.checkPersonalSex(this.common.userPersonalInformation.userSex)) {
      return this.common.sexIncorectFlag;
    }
    if (!this.common.checkPersonalBirthday(this.common.userPersonalInformation.userBirthday)) {
      return this.common.birthdayIncorectFlag;
    }
    if (!this.common.checkPhone(this.common.userPersonalInformation.userContactPhone)) {
      return this.common.phoneIncorectFlag;
    }
    if (!this.common.checkEmail(this.common.userPersonalInformation.userContactEmail)) {
      return this.common.emailFormatIncorectFlag;
    }
    return this.common.corectFlag;
  }

  changeUserSex(sex) {
    this.common.userPersonalInformation.userSex = sex;
  }

  // 请求修改用户个人信息
  askForChangingUserPersonalInformation() {
    // 前端信息格式检查
    if (this.personalUserFormIsOK() !== this.common.corectFlag) {
      this.toast.error(this.common.parseFlag(this.personalUserFormIsOK()), '提示');
      return;
    }

    // 构建数据结构
    const header = new HttpHeaders({
      'Content-Type': 'application/x-www-form-urlencoded;charset=UTF-8'
    });
    const requestHead = { headers: header };
    this.common.reqProto = {
      data: {
        userName: this.common.userPersonalInformation.userName,
        userSex: this.common.userPersonalInformation.userSex,
        userContactPhone: this.common.userPersonalInformation.userContactPhone,
        userContactEmail: this.common.userPersonalInformation.userContactEmail,
        userBirthday: this.common.userPersonalInformation.userBirthday,
      },
      orderBy: '',
      filter: '',
      page: 0,
      pageSize: 0,
    };
    this.http.post(this.common.updateUpiUrl, this.common.reqProto, requestHead).subscribe((res: any) => {
      this.common.replyProto = res;

      if (res.status !==  this.common.updateUpiSuccessFlag) {
        this.toast.warning(this.common.replyProto.msg, '提示');
        return;
      }

      this.common.storeUserPersonalInformation(this.common.userPersonalInformation);  // 更新localstore
      this.toast.success(this.common.replyProto.msg, '提示');
    });
  }
}

