(window.webpackJsonp=window.webpackJsonp||[]).push([[13],{xYuj:function(t,n,e){"use strict";e.r(n),e.d(n,"TotpModule",(function(){return l}));var o=e("ofXK"),i=e("tyNb"),s=e("PCNd"),r=e("3Pt+"),c=e("1kSV"),a=e("fXoL"),b=e("qXBG");function u(t,n){if(1&t){const t=a.Qb();a.Pb(0,"ngb-alert",17),a.ac("close",(function(){return a.qc(t),a.cc(3).ShowMessage=!1})),a.zc(1),a.Ob()}if(2&t){const t=a.cc(3);a.hc("type",t.MessageType),a.xb(1),a.Ac(t.ResponseFromBackend.Error.Message)}}function d(t,n){if(1&t&&(a.Pb(0,"div",15),a.xc(1,u,2,2,"ngb-alert",16),a.Ob()),2&t){const t=a.cc(2);a.xb(1),a.hc("ngIf",t.ShowMessage)}}function p(t,n){if(1&t){const t=a.Qb();a.Pb(0,"div",4),a.Pb(1,"div",5),a.Pb(2,"h3"),a.zc(3,"Second factor"),a.Ob(),a.Ob(),a.Pb(4,"div",6),a.Pb(5,"form",7,8),a.ac("ngSubmit",(function(){a.qc(t);const n=a.pc(6);return a.cc().OnSubmitForm(n)})),a.Pb(7,"div",9),a.Lb(8,"input",10),a.Ob(),a.Pb(9,"div",11),a.Pb(10,"div",12),a.Pb(11,"button",13),a.zc(12,"Check"),a.Ob(),a.Ob(),a.Ob(),a.Ob(),a.Ob(),a.xc(13,d,2,1,"div",14),a.Ob()}if(2&t){const t=a.pc(6),n=a.cc();a.xb(11),a.hc("disabled",t.invalid),a.xb(2),a.hc("ngIf",n.ShowMessage)}}function g(t,n){1&t&&(a.Pb(0,"div",18),a.Pb(1,"span",19),a.zc(2,"Loading..."),a.Ob(),a.Ob())}const h=[{path:"",component:(()=>{class t{constructor(t,n){this.authservice=t,this.router=n,this.LoginMode=!0,this.IsLoading=!1}ngOnDestroy(){this.SfResultSub.unsubscribe(),this.SfErrSub.unsubscribe()}ngOnInit(){this.authservice.CheckRegistered()&&this.Redirect(),this.SfErrSub=this.authservice.SfErrorSub.subscribe(t=>{switch(this.ShowMessage=!0,this.ResponseFromBackend=t,setTimeout(()=>this.ShowMessage=!1,5e3),t.Error.Code){case 200:this.MessageType="success";break;default:this.MessageType="danger"}this.IsLoading=!1}),this.SfResultSub=this.authservice.SfResultSub.subscribe(t=>{this.IsLoading=!1,t&&this.Redirect()})}OnSubmitForm(t){this.IsLoading=!0,this.authservice.SecondFactorCheck(t.value.passkey),t.reset()}Redirect(){this.router.navigate(["/recipes"])}}return t.\u0275fac=function(n){return new(n||t)(a.Kb(b.a),a.Kb(i.c))},t.\u0275cmp=a.Eb({type:t,selectors:[["app-totp"]],decls:4,vars:2,consts:[[1,"main-parent-login"],["class","card mx-auto","style","margin-top: 3px;",4,"ngIf"],[1,"text-center"],["class","spinner-border","role","status",4,"ngIf"],[1,"card","mx-auto",2,"margin-top","3px"],[1,"card-header"],[1,"card-body"],[3,"ngSubmit"],["SignupForm","ngForm"],[1,"input-group","form-group"],["type","text","name","passkey","inputmode","numeric","pattern","[0-9]*","autocomplete","one-time-code","placeholder","Passkey","ngModel","","required","","minlength","6","placement","right","ngbTooltip","Minimum 6 chars",1,"form-control"],[1,"form-group","float-right"],[1,"input-group"],["type","submit",1,"btn","btn","btn-outline-primary",3,"disabled"],["style","padding: 3px",4,"ngIf"],[2,"padding","3px"],[3,"type","close",4,"ngIf"],[3,"type","close"],["role","status",1,"spinner-border"],[1,"sr-only"]],template:function(t,n){1&t&&(a.Pb(0,"div",0),a.xc(1,p,14,2,"div",1),a.Pb(2,"div",2),a.xc(3,g,3,0,"div",3),a.Ob(),a.Ob()),2&t&&(a.xb(1),a.hc("ngIf",!n.IsLoading),a.xb(2),a.hc("ngIf",n.IsLoading))},directives:[o.l,r.r,r.i,r.j,r.b,r.n,r.h,r.k,r.o,r.e,c.k,c.a],styles:[".main-parent-login[_ngcontent-%COMP%]{position:absolute;top:50%;left:50%;transform:translate(-50%,-50%)}.card[_ngcontent-%COMP%]{width:300px}.remember[_ngcontent-%COMP%]   input[_ngcontent-%COMP%]{width:20px;height:20px;margin-left:15px;margin-right:5px}input.ng-invalid.ng-touched[_ngcontent-%COMP%]{border-color:#ff0a1d!important;box-shadow:none!important}input.ng-invalid.ng-touched[_ngcontent-%COMP%]:focus{box-shadow:0 0 0 .2rem rgba(255,10,29,.25)!important}input.ng-valid.ng-touched[_ngcontent-%COMP%]{border-color:#67c07b!important;box-shadow:none!important}input.ng-valid.ng-touched[_ngcontent-%COMP%]:focus{box-shadow:0 0 0 .2rem rgba(52,179,13,.25)!important}"]}),t})()}];let l=(()=>{class t{}return t.\u0275mod=a.Ib({type:t}),t.\u0275inj=a.Hb({factory:function(n){return new(n||t)},imports:[[o.c,r.d,i.g.forChild(h),s.a,c.b,c.l]]}),t})()}}]);