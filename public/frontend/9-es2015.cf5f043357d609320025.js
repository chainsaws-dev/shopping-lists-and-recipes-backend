(window.webpackJsonp=window.webpackJsonp||[]).push([[9],{IQtI:function(e,t,n){"use strict";n.r(t),n.d(t,"UserProfileModule",(function(){return R}));var o=n("ofXK"),r=n("AytR"),i=n("fXoL"),s=n("tyNb"),c=n("GXvH"),a=n("1kSV"),b=n("3Pt+"),d=n("mrSG"),l=n("tk/3"),u=n("qXBG");let h=(()=>{class e{constructor(e,t){this.http=e,this.auth=t}transform(e){return Object(d.a)(this,void 0,void 0,(function*(){const t=this.auth.GetUserToken(),n=new l.e({Auth:t,ApiKey:r.a.ApiKey});try{const t=yield this.http.get(e,{headers:n,responseType:"blob"}).toPromise(),o=new FileReader;return new Promise((e,n)=>{o.onloadend=()=>e(o.result),o.readAsDataURL(t)})}catch(o){return"/favicon.ico"}}))}}return e.\u0275fac=function(t){return new(t||e)(i.Mb(l.b),i.Mb(u.a))},e.\u0275pipe=i.Lb({name:"secureImagePipe",type:e,pure:!0}),e})();function p(e,t){1&e&&(i.Rb(0,"div",5),i.Rb(1,"span",6),i.Bc(2,"Loading..."),i.Qb(),i.Qb())}function g(e,t){if(1&e){const e=i.Sb();i.Rb(0,"ngb-alert",7),i.cc("close",(function(){return i.sc(e),i.ec().ShowMessage=!1})),i.Bc(1),i.Qb()}if(2&e){const e=i.ec();i.jc("type",e.MessageType),i.zb(1),i.Dc(" ",e.ResponseFromBackend.Error.Message,"")}}function m(e,t){if(1&e){const e=i.Sb();i.Rb(0,"form",8,9),i.cc("ngSubmit",(function(){i.sc(e);const t=i.rc(1);return i.ec().OnSaveClick(t)})),i.Rb(2,"div",10),i.Rb(3,"div",11),i.Rb(4,"h2"),i.Bc(5,"Profile"),i.Qb(),i.Rb(6,"div",10),i.Rb(7,"label",12),i.Bc(8,"Name"),i.Qb(),i.Nb(9,"input",13),i.Qb(),i.Rb(10,"div",10),i.Rb(11,"label",14),i.Bc(12,"Email"),i.Qb(),i.Nb(13,"input",15),i.Qb(),i.Rb(14,"div",10),i.Rb(15,"label",16),i.Bc(16,"Phone"),i.Qb(),i.Nb(17,"input",17),i.Qb(),i.Rb(18,"label",18),i.Bc(19,"Password"),i.Qb(),i.Rb(20,"div",19),i.Rb(21,"div",20),i.Rb(22,"div",21),i.Rb(23,"input",22),i.cc("ngModelChange",(function(t){return i.sc(e),i.ec().changepassword=t})),i.Qb(),i.Qb(),i.Qb(),i.Nb(24,"input",23),i.Qb(),i.Rb(25,"button",24),i.Bc(26,"Save"),i.Qb(),i.Rb(27,"button",25),i.Bc(28,"Cancel"),i.Qb(),i.Qb(),i.Qb(),i.Qb()}if(2&e){const e=i.rc(1),t=i.ec();i.zb(9),i.jc("ngModel",t.UserToEdit.Name),i.zb(4),i.jc("ngModel",t.UserToEdit.Email),i.zb(4),i.jc("ngModel",t.UserToEdit.Phone),i.zb(6),i.jc("ngModel",t.changepassword),i.zb(1),i.jc("disabled",!t.changepassword),i.zb(1),i.jc("disabled",e.invalid)}}function f(e,t){if(1&e){const e=i.Sb();i.Rb(0,"div",10),i.Rb(1,"h2"),i.Bc(2,"Second factor enabled"),i.Qb(),i.Rb(3,"button",26),i.cc("click",(function(){return i.sc(e),i.ec().OnUnlinkTwoFactor()})),i.Bc(4,"Disable"),i.Qb(),i.Qb()}}function w(e,t){if(1&e){const e=i.Sb();i.Rb(0,"form",8,27),i.cc("ngSubmit",(function(){i.sc(e);const t=i.rc(1);return i.ec().OnLinkTwoFactor(t)})),i.Rb(2,"div",10),i.Rb(3,"div",11),i.Rb(4,"h2"),i.Bc(5,"Set second factor"),i.Qb(),i.Rb(6,"h4"),i.Bc(7,"1. Scan with "),i.Rb(8,"a",28),i.Bc(9,"Authenticator"),i.Qb(),i.Bc(10,": "),i.Qb(),i.Nb(11,"img",29),i.fc(12,"async"),i.fc(13,"secureImagePipe"),i.Rb(14,"h4"),i.Bc(15,"2. Enter the token:"),i.Qb(),i.Rb(16,"label",18),i.Bc(17,"Token"),i.Qb(),i.Rb(18,"div",19),i.Rb(19,"div",20),i.Rb(20,"div",21),i.Rb(21,"input",30),i.cc("ngModelChange",(function(t){return i.sc(e),i.ec().SetTwoFactor=t})),i.Qb(),i.Qb(),i.Qb(),i.Nb(22,"input",31),i.Qb(),i.Rb(23,"button",32),i.Bc(24,"Save"),i.Qb(),i.Qb(),i.Qb(),i.Qb()}if(2&e){const e=i.ec();i.zb(8),i.jc("href",e.AuthUrl,i.tc),i.zb(3),i.jc("src",i.gc(12,4,i.gc(13,6,e.QrUrl)),i.tc),i.zb(10),i.jc("ngModel",e.SetTwoFactor),i.zb(1),i.jc("disabled",!e.SetTwoFactor)}}const v=[{path:"",component:(()=>{class e{constructor(e,t,n){this.activatedroute=e,this.router=t,this.datastore=n}ngOnDestroy(){this.FetchUser.unsubscribe(),this.RecivedErrorSub.unsubscribe(),this.DataLoading.unsubscribe(),this.SaveUser.unsubscribe(),this.LinkUnlinkTFA.unsubscribe()}ngOnInit(){this.AuthUrl=r.a.GetAuthenticatorUrl,this.QrUrl=r.a.GetTOTPQRCodeUrl,this.LinkUnlinkTFA=this.datastore.TwoFactorSub.subscribe(e=>{this.SetUserAndTFA(e)}),this.SaveUser=this.datastore.UserUpdateInsert.subscribe(e=>{this.SetUserAndTFA(e)}),this.FetchUser=this.datastore.CurrentUserFetch.subscribe(e=>{this.SetUserAndTFA(e)}),this.RecivedErrorSub=this.datastore.RecivedError.subscribe(e=>{if(this.ShowMessage=!0,this.ResponseFromBackend=e,setTimeout(()=>this.ShowMessage=!1,5e3),e)switch(e.Error.Code){case 200:this.MessageType="success";break;default:this.MessageType="danger"}}),this.DataLoading=this.datastore.LoadingData.subscribe(e=>{this.IsLoading=e}),this.datastore.FetchCurrentUser()}SetUserAndTFA(e){this.UserToEdit=e,this.UserToEdit&&(this.TwoFactorEnabled=this.UserToEdit.SecondFactor)}OnSaveClick(e){if(e.valid){if(e.value.changepassword&&0===e.value.newpassword.length)return;this.UserToEdit.Email=e.value.useremail,this.UserToEdit.Name=e.value.username,this.UserToEdit.Phone=e.value.userphone,this.datastore.SaveCurrentUser(this.UserToEdit,e.value.changepassword,e.value.newpassword)}}OnLinkTwoFactor(e){e.valid&&this.datastore.LinkTwoFactor(e.value.passkey,this.UserToEdit)}OnUnlinkTwoFactor(){this.UserToEdit.SecondFactor&&this.datastore.UnlinkTwoFactor(this.UserToEdit)}}return e.\u0275fac=function(t){return new(t||e)(i.Mb(s.a),i.Mb(s.c),i.Mb(c.a))},e.\u0275cmp=i.Gb({type:e,selectors:[["app-user-profile"]],decls:6,vars:5,consts:[[1,"text-center"],["class","spinner-border","role","status",4,"ngIf"],[3,"type","close",4,"ngIf"],[3,"ngSubmit",4,"ngIf"],["class","form-group",4,"ngIf"],["role","status",1,"spinner-border"],[1,"sr-only"],[3,"type","close"],[3,"ngSubmit"],["UserProfileForm","ngForm"],[1,"form-group"],[2,"margin","3px"],["for","name"],["type","text","id","name","placeholder","example","name","username","required","",1,"form-control","mb-1",3,"ngModel"],["for","email"],["type","email","id","email","placeholder","exampe@example.com","name","useremail","required","","email","",1,"form-control","mb-1",3,"ngModel"],["for","phone"],["type","tel","id","phone","placeholder","+7 (965) 777-77-77","name","userphone","pattern","^((8|\\+7)[\\- ]?)?(\\(?\\d{3,4}\\)?[\\- ]?)?[\\d\\- ]{5,10}$",1,"form-control","mb-1",3,"ngModel"],["for","newpassword"],[1,"input-group","mb-3"],[1,"input-group-prepend"],[1,"input-group-text"],["type","checkbox","id","changepassword","name","changepassword",3,"ngModel","ngModelChange"],["type","password","id","newpassword","name","newpassword","placeholder","Enter new password","ngModel","",1,"form-control",3,"disabled"],["type","submit",1,"btn","btn-outline-primary",3,"disabled"],["type","button","routerLink","..",1,"btn","btn-outline-danger",2,"margin-left","3px"],["type","button",1,"btn","btn-outline-danger",3,"click"],["UserTwoFactorForm","ngForm"],[3,"href"],["alt","QR code",3,"src"],["type","checkbox","id","enabletwofactor","name","enabletwofactor",3,"ngModel","ngModelChange"],["type","text","name","passkey","inputmode","numeric","pattern","[0-9]*","autocomplete","one-time-code","id","newpasskey","placeholder","Token from Authenticator","ngModel","","required","","minlength","6","ngbTooltip","Minimum 6 numbers",1,"form-control",3,"disabled"],["type","submit",1,"btn","btn-outline-primary"]],template:function(e,t){1&e&&(i.Rb(0,"div",0),i.zc(1,p,3,0,"div",1),i.Qb(),i.zc(2,g,2,2,"ngb-alert",2),i.zc(3,m,29,6,"form",3),i.zc(4,f,5,0,"div",4),i.zc(5,w,25,8,"form",3)),2&e&&(i.zb(1),i.jc("ngIf",t.IsLoading),i.zb(1),i.jc("ngIf",t.ShowMessage),i.zb(1),i.jc("ngIf",t.UserToEdit),i.zb(1),i.jc("ngIf",t.UserToEdit&&t.TwoFactorEnabled),i.zb(1),i.jc("ngIf",t.UserToEdit&&!t.TwoFactorEnabled))},directives:[o.l,a.a,b.s,b.j,b.k,b.c,b.p,b.i,b.l,b.d,b.o,b.a,s.d,b.f,a.k],pipes:[o.b,h],styles:["input.ng-invalid.ng-touched[_ngcontent-%COMP%]{border-color:#ff0a1d!important;box-shadow:none!important}input.ng-invalid.ng-touched[_ngcontent-%COMP%]:focus{box-shadow:0 0 0 .2rem rgba(255,10,29,.25)!important}input.ng-valid.ng-touched[_ngcontent-%COMP%]{border-color:#67c07b!important;box-shadow:none!important}input.ng-valid.ng-touched[_ngcontent-%COMP%]:focus{box-shadow:0 0 0 .2rem rgba(52,179,13,.25)!important}select.ng-invalid.ng-touched[_ngcontent-%COMP%]{border-color:#ff0a1d!important;box-shadow:none!important}select.ng-invalid.ng-touched[_ngcontent-%COMP%]:focus{box-shadow:0 0 0 .2rem rgba(255,10,29,.25)!important}select.ng-valid.ng-touched[_ngcontent-%COMP%]{border-color:#67c07b!important;box-shadow:none!important}select.ng-valid.ng-touched[_ngcontent-%COMP%]:focus{box-shadow:0 0 0 .2rem rgba(52,179,13,.25)!important}.image-placeholder[_ngcontent-%COMP%]{background-color:#eee;display:flex;height:333px;margin:5px;border-radius:3px;max-width:592px}.image-placeholder[_ngcontent-%COMP%] > h4[_ngcontent-%COMP%]{align-self:center;text-align:center;width:100%}"]}),e})()}];let R=(()=>{class e{}return e.\u0275mod=i.Kb({type:e}),e.\u0275inj=i.Jb({factory:function(t){return new(t||e)},imports:[[o.c,b.e,a.b,s.g.forChild(v),a.l]]}),e})()},mrSG:function(e,t,n){"use strict";function o(e,t,n,o){return new(n||(n=Promise))((function(r,i){function s(e){try{a(o.next(e))}catch(t){i(t)}}function c(e){try{a(o.throw(e))}catch(t){i(t)}}function a(e){var t;e.done?r(e.value):(t=e.value,t instanceof n?t:new n((function(e){e(t)}))).then(s,c)}a((o=o.apply(e,t||[])).next())}))}n.d(t,"a",(function(){return o}))}}]);