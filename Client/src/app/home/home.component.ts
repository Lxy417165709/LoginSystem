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
    private toast: ToastrService,
  ) { }

  ngOnInit() {
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

  // 退出
  exit() {
    this.common.storeUserPersonalInformation(null);
    this.common.storeUserAccountInformation(null);
    this.toast.success('退出成功!', '提示', {positionClass: 'toast-bottom-right'});
    this.router.navigate(['login']);
  }
}
