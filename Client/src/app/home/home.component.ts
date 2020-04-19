import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { CommonService } from '../common.service';
import { DomSanitizer } from '@angular/platform-browser';
import { ToastrService } from 'ngx-toastr';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {


  constructor(
    private http: HttpClient,
    private router: Router,
    private common: CommonService,
    private sanitizer: DomSanitizer,
    private toast: ToastrService,
  ) { }


  getUpi() {
    const header = new HttpHeaders({
      'Content-Type': 'application/x-www-form-urlencoded;charset=UTF-8'
    });

    const requestHead = { headers: header };

    // 通过token 获取用户信息
    this.http.post('/server/getUpi', this.common.reqProto, requestHead).subscribe((res: any) => {
      console.log(res);
      this.common.replyProto = res;

      // 根据登录结果相应操作
      // 状态码为0表示失败
      if (res.status === 0) {
        this.toast.warning(res.msg, '提示');
        this.router.navigate(['login']);
        return;
      }

      // 存储用户账户信息
      this.common.storeUserPersonalInformation(this.common.replyProto.data);
      this.common.userPersonalInformation = this.common.getUserPersonalInformation();
      console.log(this.common.userPersonalInformation);
      this.toast.success('个人信息获取成功!', '提示', {positionClass: 'toast-bottom-right'});
    });
  }
  getUai() {

    const header = new HttpHeaders({
      'Content-Type': 'application/x-www-form-urlencoded;charset=UTF-8'
    });

    const requestHead = { headers: header };

    // 通过token 获取用户信息
    this.http.post('/server/getUai', this.common.reqProto, requestHead).subscribe((res: any) => {
      this.common.replyProto = res;
      console.log(res);
      // 根据登录结果相应操作
        // 状态码为0表示失败
      if (res.status === 0) {
        // 登录失败时的提示
        this.toast.warning(res.msg, '提示');
        this.router.navigate(['login']);
        return;
      }

      // 存储用户账户信息
      this.common.storeUserAccountInformation(this.common.replyProto.data);
      this.toast.success('账号信息获取成功!', '提示', {positionClass: 'toast-bottom-right'});

      // // 检测登录状态
      // if (this.common.loginStateDetection() === false) {
      //   this.toast.error('您还没登录', '提示');
      //   this.router.navigate(['login']);
      //   return;
      // }
    });
  }



  ngOnInit() {
    this.getUai();
    this.getUpi();
  }

  // 退出
  exit() {
    this.toast.success('退出成功!', '提示', {positionClass: 'toast-bottom-right'});
    this.router.navigate(['login']);
  }


  // 用这种没有返回值的获取用户个人信息吧
  askForPersonalInformation() {
    const header = new HttpHeaders({
      'Content-Type': 'application/x-www-form-urlencoded;charset=UTF-8'
    });

    // 请求头
    const requestHead = { headers: header };

    // 请求协议 (请求体)
    this.common.reqProto = {
      data: {
        userId: this.common.userAccountInformation.userId // 请求数据只需要用户id
      },
      // ---- 下面的字段都没用到
      orderBy: '',  // 排序要求
      filter: '',   // 筛选条件
      page: 0,      // 分页
      pageSize: 0,  // 分页大小
    };

    this.http.post('/server/getPersonalInformation', this.common.reqProto, requestHead).subscribe((res: any) => {

      // 返回逻辑还有很多没考虑
      this.common.replyProto = res;

      // 根据返回状态执行相应操作 (0 表示成功)
      if (this.common.replyProto.status === 0) {

        // 用户的头像base64
        this.common.userPersonalInformation = res.data.UserPersonalInformation;

        // 获取dataUrl
        const dataURL = 'data:image/jpeg;base64,' + res.data.UserPhotoData ;

        // 获取blobURL对象
        const blobURLObject = this.sanitizer.bypassSecurityTrustUrl(URL.createObjectURL(this.common.dataURLtoBlob(dataURL)));

        // 获取blobUrl字符串并构成安全链接(不安全会报错)
        this.common.userPersonalInformation.userPhotoUrl = blobURLObject;

        // 存储用户个人信息到sessionStorage (成功了)
        sessionStorage.setItem('userPersonalInformation', JSON.stringify(this.common.userPersonalInformation));

      } else {
        // 这里表示出错
        this.toast.error(res.msg, '提示');
        this.router.navigate(['login']);
      }
      // console.log(this.common.replyProto);
    });
  }
}
