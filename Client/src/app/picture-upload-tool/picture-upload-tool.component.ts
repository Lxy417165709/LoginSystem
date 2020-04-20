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

  // 一个传值器，用于通知父组件关闭该子组件
  @Output() emitter = new EventEmitter();


  // 头像上传器
  uploader: FileUploader = new FileUploader({
    url: '/server/uploadPhoto',
    method: 'POST',
    itemAlias: 'file'
  });

  constructor(
    private common: CommonService,
    private sanitizer: DomSanitizer,
    private toast: ToastrService,
    private http: HttpClient,
  ) { }

  ngOnInit() {
    // this.userPersonalInformation 在这里现在好像没啥用
    this.common.userAccountInformation = this.common.getUserAccountInformation();
    this.common.userPersonalInformation = this.common.getUserPersonalInformation();
    this.tmpUrl = this.getImgSrcPhoto();
    // this.common.userPersonalInformation.photoData;
  }
  getImgSrcPhoto() {
    return 'data:image/jpg;base64,' + this.common.userPersonalInformation.photoData;
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
    this.common.userPersonalInformation.userPhotoUrl = this.tmpUrl;
    this.emitter.emit({
      type: 1,
      data: this.common.userPersonalInformation,
    });
  }

  // 选择图片发生错误时执行的操作提示
  // 第一个参数是提示语，第二个是图片选择器主体
  selectError(tipStr, input) {
    this.tmpUrl =  this.common.userPersonalInformation.userPhotoUrl;
    this.toast.error(tipStr, '提示');
    input.value = null;
    this.uploader.queue = [];
  }

  // 选择图片后的操作
  listenFileChanged(e) {
    // 当上传器有照片时，我们要去除第一个
    if (this.uploader.queue.length === 2) {
      const wannaElement = this.uploader.queue[1];
      this.uploader.queue = [];
      this.uploader.queue.push(wannaElement);
    }
    if (e.target.files[0] === undefined) {
      this.selectError('您还没有选择头像!', e.target);
      return ;
    }
     // 判断是否是图片
    const isPhoto = this.common.isPhotoFile(e.target.files[0]);
    if (!isPhoto) {
      this.selectError('您选择的文件不是一个图片文件!', e.target);
      return ;
    }
     // 判断图片文件大小,限制图片大小只能在500KB内
    const size = e.target.files[0].size;
    if (size > 0.6 * 1024 * 1024) {
      this.selectError('图片大小只能在500KB(0.5MB)范围内!', e.target);
      this.tmpUrl = this.getImgSrcPhoto();
      return ;
    }
    const file = e.target.files[0];
    const reader = new FileReader();
    reader.readAsDataURL(file);
    reader.onloadend = (en) => {
      // console.log(en.target.result);
      // this.tmpUrl = (en.target as any).result.toString();

      this.tmpUrl = (en.target as any).result.toString();
    };
    this.toast.success('新头像载入成功!', '提示');
  }

  // 上传图片 (成功了!)
  // 该函数的作用是改变用户头像
  askForChangingUserPhoto() {

    // 用户没有文件上传时
    if (this.uploader.queue[0] === undefined) {
      this.toast.error('您还没有选择任何新头像!', '提示');
      return;
    }

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


    // 发送注册请求
    this.http.post('/server/updatePhoto', this.common.reqProto, requestHead).subscribe((res: any) => {
      // 成功了！ 但是这的业务逻辑还有好多
      this.common.replyProto = res;
      console.log(res);
      // 状态码为0表示失败
      if (res.status === 0) {
        // 输出响应信息字段
        this.toast.warning(res.msg, '提示');
        return;
      }

      this.toast.success('头像修改成功!', '提示', {timeOut: 4000});
      this.common.userPersonalInformation.photoData = this.common.getPureBase64(this.tmpUrl);
    });

    // // 成功回调
    // this.uploader.queue[0].onSuccess = (res: any) => {
    //   this.common.replyProto = JSON.parse(res); // 把返回结果存在this.common.replyProto中

    //   // 这里表示成功了
    //   if (this.common.replyProto.status === 0) {
    //     this.toast.success('头像上传成功,即将更改个人信息!', '提示');
    //     // 修改头像名
    //     this.common.userPersonalInformation.userPhoto = this.uploader.queue[0].file.name;

    //     // 清空队列
    //     this.uploader.queue = [];

    //     // 通知父组件修改用户个人信息
    //     this.callFatherExecChange();

    //     // 通知父组件关闭该子组件
    //     this.callFatherExecClose();
    //   } else {
    //     // 上传失败提醒
    //     this.toast.warning('头像上传失败, 请重试!', '提示');
    //   }
    // };
  }
}
