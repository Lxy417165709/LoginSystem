import { Component, OnInit, Output } from '@angular/core';
import { CommonService } from '../common.service';
import { FileUploader } from 'ng2-file-upload';
import { EventEmitter } from '@angular/core';
import { DomSanitizer } from '@angular/platform-browser';
import { ToastrService } from 'ngx-toastr';
import { HttpClient, HttpHeaders } from '@angular/common/http';

@Component({
  selector: 'app-picture-upload-tool',
  templateUrl: './picture-upload-tool.component.html',
  styleUrls: ['./picture-upload-tool.component.css']
})
export class PictureUploadToolComponent implements OnInit {

  // 临时url,用于展示用户选择的图片
  tmpUrl: string;
  loadSuccessMsg = '新头像载入成功!';
  // 一个传值器，用于通知父组件关闭该子组件
  @Output() emitter = new EventEmitter();


  constructor(
    private common: CommonService,
    private toast: ToastrService,
    private http: HttpClient,
  ) { }

  ngOnInit() {
    this.common.userAccountInformation = this.common.getUserAccountInformation();
    this.common.userPersonalInformation = this.common.getUserPersonalInformation();
    this.tmpUrl = this.common.getUserPhotoUrl();
  }

  // 向父组件传值，通知父组件关闭该子组件 (传送0: 表示关闭该子组件)
  callFatherExecClose() {
    this.emitter.emit({
      type: 0,
      data: 0,
    });
  }

  // 向父组件传值，通知父组件修改用户个人信息
  callFatherExecChange() {
    this.emitter.emit({
      type: 1,
      data: this.common.userPersonalInformation,
    });
  }


  selectCheck(file) {
    if (file === undefined) {
      return this.common.photoEmptyFlag;
    }
    if (!this.common.isPhotoFile(file)) {
      return this.common.notPhotoFlag;
    }
    if (file.size > this.common.photoMaxSize) {
      return this.common.photoSizeTooBigFlag;
    }
    return this.common.corectFlag;
  }

  photoCheck(photo) {
    if (this.tmpUrl === this.common.getUserPhotoUrl()) {
      return this.common.notNewPhotoFlag;
    }
    return this.common.corectFlag;
  }
  // 监听图片选择
  listenFileChanged(e) {
    if (this.selectCheck(e.target.files[0]) !== this.common.corectFlag) {
      this.toast.error(this.common.parseFlag(this.selectCheck(e.target.files[0])), '提示');
      e.value = null;
      return;
    }

    const reader = new FileReader();
    reader.readAsDataURL(e.target.files[0]);
    reader.onloadend = (en) => {
      this.tmpUrl = (en.target as any).result.toString();
    };
    this.toast.success(this.loadSuccessMsg, '提示');
  }

  // 该函数的作用是改变用户头像
  askForChangingUserPhoto() {

    // 校验
    if (this.photoCheck(this.tmpUrl) !== this.common.corectFlag) {
      this.toast.error(this.common.parseFlag(this.photoCheck(this.tmpUrl)), '提示');
      return;
    }

    // 数据结构构建
    const header = new HttpHeaders({
      'Content-Type': 'application/x-www-form-urlencoded;charset=UTF-8'
    });
    const requestHead = { headers: header };

    this.common.reqProto = {
      data: {
        photoBase64: this.common.getPureBase64(this.tmpUrl),
      },   // 请求数据
      orderBy: '',  // 排序要求
      filter: '',   // 筛选条件
      page: 0,      // 分页
      pageSize: 0,  // 分页大小
    };


    // 发送更新头像请求
    this.http.post(this.common.updatePhotoUrl, this.common.reqProto, requestHead).subscribe((res: any) => {
      this.common.replyProto = res;

      if (this.common.replyProto.status !== this.common.updatePhotoSuccessFlag) {
        this.toast.warning(this.common.replyProto.msg, '提示');
        return;
      }

      this.common.userPersonalInformation.photoData = this.common.getPureBase64(this.tmpUrl);
      this.common.storeUserPersonalInformation(this.common.userPersonalInformation);
      this.toast.success(this.common.replyProto.msg, '提示', {timeOut: 4000});
      // this.callFatherExecChange();
      // this.callFatherExecClose();
    });
  }
}
