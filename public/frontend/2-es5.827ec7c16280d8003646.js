!function(){function e(e,n){if(!(e instanceof n))throw new TypeError("Cannot call a class as a function")}function n(e,n){for(var t=0;t<n.length;t++){var o=n[t];o.enumerable=o.enumerable||!1,o.configurable=!0,"value"in o&&(o.writable=!0),Object.defineProperty(e,o.key,o)}}(window.webpackJsonp=window.webpackJsonp||[]).push([[2],{grO1:function(t,o,r){"use strict";r.r(o),r.d(o,"ConfirmEmailModule",(function(){return M}));var i=r("ofXK"),a=r("tyNb"),s=r("fXoL"),c=r("GXvH"),d=r("3Pt+"),u=r("1kSV");function b(e,n){1&e&&s.Lb(0,"input",18)}function l(e,n){1&e&&s.Lb(0,"input",19)}function f(e,n){if(1&e&&(s.Pb(0,"div",20),s.Pb(1,"div",21),s.Pb(2,"button",22),s.wc(3),s.Ob(),s.Ob(),s.Ob()),2&e){s.cc();var t=s.mc(2),o=s.cc(2);s.xb(2),s.fc("disabled",t.invalid),s.xb(1),s.xc(o.ResetPasswordMode&&o.Token?"Save":"Send")}}function g(e,n){if(1&e){var t=s.Qb();s.Pb(0,"div",11),s.Pb(1,"form",12,13),s.ac("ngSubmit",(function(){s.nc(t);var e=s.mc(2);return s.cc(2).OnSubmitForm(e)})),s.Pb(3,"div",14),s.uc(4,b,1,0,"input",15),s.uc(5,l,1,0,"input",16),s.Ob(),s.uc(6,f,4,2,"div",17),s.Ob(),s.Ob()}if(2&e){var o=s.cc(2);s.xb(4),s.fc("ngIf",!o.Token),s.xb(1),s.fc("ngIf",o.ResetPasswordMode&&o.Token),s.xb(1),s.fc("ngIf",!o.Token||o.ResetPasswordMode)}}function p(e,n){if(1&e){var t=s.Qb();s.Pb(0,"ngb-alert",25),s.ac("close",(function(){return s.nc(t),s.cc(3).ShowMessage=!1})),s.wc(1),s.Ob()}if(2&e){var o=s.cc(3);s.fc("type",o.MessageType),s.xb(1),s.xc(o.ResponseFromBackend.Error.Message)}}function m(e,n){if(1&e&&(s.Pb(0,"div",23),s.uc(1,p,2,2,"ngb-alert",24),s.Ob()),2&e){var t=s.cc(2);s.xb(1),s.fc("ngIf",t.ShowMessage)}}function h(e,n){if(1&e&&(s.Pb(0,"div",4),s.Pb(1,"div",5),s.Pb(2,"h3"),s.wc(3),s.Ob(),s.Ob(),s.uc(4,g,7,3,"div",6),s.Pb(5,"div",7),s.Pb(6,"div",8),s.Pb(7,"a",9),s.wc(8,"Back"),s.Ob(),s.Ob(),s.Ob(),s.uc(9,m,2,1,"div",10),s.Ob()),2&e){var t=s.cc();s.xb(3),s.xc(t.ResetPasswordMode?"Reset password":"Confirm email"),s.xb(1),s.fc("ngIf",!t.Token||t.ResetPasswordMode),s.xb(5),s.fc("ngIf",t.ShowMessage)}}function v(e,n){1&e&&(s.Pb(0,"div",26),s.Pb(1,"span",27),s.wc(2,"Loading..."),s.Ob(),s.Ob())}var w,P,y=[{path:"",component:(w=function(){function t(n,o,r){e(this,t),this.DataServ=n,this.activeroute=o,this.router=r,this.IsLoading=!1}var o,r,i;return o=t,(r=[{key:"ngOnDestroy",value:function(){this.RecivedErrorSub.unsubscribe(),this.RecivedResponseSub.unsubscribe(),this.DataServiceSub.unsubscribe()}},{key:"ngOnInit",value:function(){var e=this;this.DataServiceSub=this.DataServ.LoadingData.subscribe((function(n){e.IsLoading=n})),this.RecivedErrorSub=this.DataServ.RecivedError.subscribe((function(n){switch(e.ShowMessage=!0,e.ResponseFromBackend=n,setTimeout((function(){return e.ShowMessage=!1}),5e3),n.Error.Code){case 200:e.MessageType="success";break;default:e.MessageType="danger"}})),this.activeroute.queryParams.subscribe((function(n){e.Token=n.Token;var t=e.getUrlWithoutParams();e.ResetPasswordMode="/reset-password"===t,e.ResetPasswordMode||e.Token&&e.DataServ.ConfirmEmail(e.Token)}))}},{key:"getUrlWithoutParams",value:function(){var e=this.router.parseUrl(this.router.url);return e.queryParams={},e.toString()}},{key:"OnSubmitForm",value:function(e){e.valid&&(this.IsLoading=!0,this.ResetPasswordMode&&this.Token?this.DataServ.SubmitNewPassword(this.Token,e.value.newpassword):this.ResetPasswordMode?this.DataServ.SendEmailResetPassword(e.value.email):this.DataServ.SendEmailConfirmEmail(e.value.email),e.reset())}}])&&n(o.prototype,r),i&&n(o,i),t}(),w.\u0275fac=function(e){return new(e||w)(s.Kb(c.a),s.Kb(a.a),s.Kb(a.c))},w.\u0275cmp=s.Eb({type:w,selectors:[["app-confirm-email"]],decls:4,vars:2,consts:[[1,"main-parent-login"],["class","card mx-auto","style","margin-top: 3px;",4,"ngIf"],[1,"text-center"],["class","spinner-border","role","status",4,"ngIf"],[1,"card","mx-auto",2,"margin-top","3px"],[1,"card-header"],["class","card-body",4,"ngIf"],[1,"card-footer"],[1,"d-flex","justify-content-center"],["href","/"],["style","padding: 3px",4,"ngIf"],[1,"card-body"],[3,"ngSubmit"],["ResendConfEmailForm","ngForm"],[1,"input-group","form-group"],["type","email","name","email","class","form-control","placeholder","E-mail","placement","right","ngbTooltip","Registered email required","ngModel","","email","","required","",4,"ngIf"],["type","password","name","newpassword","class","form-control","placeholder","New password","ngModel","","minlength","6","placement","right","ngbTooltip","Minimum 6 chars","required","",4,"ngIf"],["class","form-group float-right",4,"ngIf"],["type","email","name","email","placeholder","E-mail","placement","right","ngbTooltip","Registered email required","ngModel","","email","","required","",1,"form-control"],["type","password","name","newpassword","placeholder","New password","ngModel","","minlength","6","placement","right","ngbTooltip","Minimum 6 chars","required","",1,"form-control"],[1,"form-group","float-right"],[1,"input-group"],["type","submit",1,"btn","btn","btn-outline-primary",3,"disabled"],[2,"padding","3px"],[3,"type","close",4,"ngIf"],[3,"type","close"],["role","status",1,"spinner-border"],[1,"sr-only"]],template:function(e,n){1&e&&(s.Pb(0,"div",0),s.uc(1,h,10,3,"div",1),s.Pb(2,"div",2),s.uc(3,v,3,0,"div",3),s.Ob(),s.Ob()),2&e&&(s.xb(1),s.fc("ngIf",!n.IsLoading),s.xb(2),s.fc("ngIf",n.IsLoading))},directives:[i.k,d.r,d.i,d.j,d.b,u.k,d.h,d.k,d.c,d.o,d.e,u.a],styles:[".main-parent-login[_ngcontent-%COMP%]{position:absolute;top:50%;left:50%;transform:translate(-50%,-50%)}.card[_ngcontent-%COMP%]{width:300px}.remember[_ngcontent-%COMP%]   input[_ngcontent-%COMP%]{width:20px;height:20px;margin-left:15px;margin-right:5px}input.ng-invalid.ng-touched[_ngcontent-%COMP%]{border-color:#ff0a1d!important;box-shadow:none!important}input.ng-invalid.ng-touched[_ngcontent-%COMP%]:focus{box-shadow:0 0 0 .2rem rgba(255,10,29,.25)!important}input.ng-valid.ng-touched[_ngcontent-%COMP%]{border-color:#67c07b!important;box-shadow:none!important}input.ng-valid.ng-touched[_ngcontent-%COMP%]:focus{box-shadow:0 0 0 .2rem rgba(52,179,13,.25)!important}"]}),w)}],M=((P=function n(){e(this,n)}).\u0275mod=s.Ib({type:P}),P.\u0275inj=s.Hb({factory:function(e){return new(e||P)},imports:[[i.b,d.d,u.b,a.g.forChild(y),u.l]]}),P)}}])}();