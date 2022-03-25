"use strict";(self.webpackChunkshopping_lists_and_recipes=self.webpackChunkshopping_lists_and_recipes||[]).push([[864],{7864:(Q,f,a)=>{a.r(f),a.d(f,{UserProfileModule:()=>N,httpTranslateLoader:()=>x});var u=a(1048),c=a(2340),e=a(9724),T=a(384),b=a(3649),Z=a(4800),p=a(4190),g=a(2719),s=a(1659),U=a(1404),v=a(8177),_=a(4762),h=a(9725),A=a(7472);let L=(()=>{class n{constructor(t,o){this.http=t,this.auth=o}transform(t){return(0,_.mG)(this,void 0,void 0,function*(){const o=this.auth.GetUserToken(),r=new h.WM({Auth:o,ApiKey:c.N.ApiKey});try{const l=yield function w(n,i){const t="object"==typeof i;return new Promise((o,r)=>{let d,l=!1;n.subscribe({next:m=>{d=m,l=!0},error:r,complete:()=>{l?o(d):t?o(i.defaultValue):r(new A.K)}})})}(this.http.get(t,{headers:r,responseType:"blob"})),d=new FileReader;return new Promise((m,I)=>{d.onloadend=()=>m(d.result),d.readAsDataURL(l)})}catch(l){return"/favicon.ico"}})}}return n.\u0275fac=function(t){return new(t||n)(e.Y36(h.eN,16),e.Y36(T.e,16))},n.\u0275pipe=e.Yjl({name:"secureImagePipe",type:n,pure:!0}),n})();function C(n,i){1&n&&(e.TgZ(0,"div",5)(1,"span",6),e._uU(2),e.ALo(3,"translate"),e.qZA()()),2&n&&(e.xp6(2),e.Oqu(e.lcZ(3,1,"Loading")))}function M(n,i){if(1&n){const t=e.EpF();e.TgZ(0,"ngb-alert",7),e.NdJ("close",function(){return e.CHM(t),e.oxw().ShowMessage=!1}),e._uU(1),e.ALo(2,"FirstLetterUpperPipe"),e.qZA()}if(2&n){const t=e.oxw();e.Q6J("type",t.MessageType),e.xp6(1),e.hij(" ",e.lcZ(2,2,t.ResponseFromBackend.Error.Message),"")}}function P(n,i){if(1&n&&(e.TgZ(0,"option",29),e._uU(1),e.qZA()),2&n){const t=i.$implicit,o=e.oxw(2);e.Q6J("value",t)("selected",t===o.translate.currentLang),e.xp6(1),e.hij(" ",t," ")}}function E(n,i){if(1&n){const t=e.EpF();e.TgZ(0,"form",8,9),e.NdJ("ngSubmit",function(){e.CHM(t);const r=e.MAs(1);return e.oxw().OnSaveClick(r)}),e.TgZ(2,"div",10)(3,"div",11)(4,"h2"),e._uU(5),e.ALo(6,"translate"),e.qZA(),e.TgZ(7,"div",10)(8,"label",12),e._uU(9),e.ALo(10,"translate"),e.qZA(),e._UZ(11,"input",13),e.ALo(12,"translate"),e.qZA(),e.TgZ(13,"div",10)(14,"label",14),e._uU(15),e.ALo(16,"translate"),e.qZA(),e._UZ(17,"input",15),e.qZA(),e.TgZ(18,"div",10)(19,"label",16),e._uU(20),e.ALo(21,"translate"),e.qZA(),e._UZ(22,"input",17),e.qZA(),e.TgZ(23,"div",10)(24,"label",18),e._uU(25),e.ALo(26,"translate"),e.qZA(),e.TgZ(27,"select",19,20),e.NdJ("change",function(){e.CHM(t);const r=e.MAs(28);return e.oxw().SwitchLanguage(r.value)}),e.ALo(29,"translate"),e.YNc(30,P,2,3,"option",21),e.qZA()(),e.TgZ(31,"label",22),e._uU(32),e.ALo(33,"translate"),e.qZA(),e.TgZ(34,"div",23)(35,"div",24)(36,"input",25),e.NdJ("ngModelChange",function(r){return e.CHM(t),e.oxw().changepassword=r}),e.qZA()(),e._UZ(37,"input",26),e.ALo(38,"translate"),e.qZA(),e.TgZ(39,"button",27),e._uU(40),e.ALo(41,"translate"),e.qZA(),e.TgZ(42,"button",28),e._uU(43),e.ALo(44,"translate"),e.qZA()()()()}if(2&n){const t=e.MAs(1),o=e.oxw();e.xp6(5),e.Oqu(e.lcZ(6,19,"Profile")),e.xp6(4),e.Oqu(e.lcZ(10,21,"Name")),e.xp6(2),e.Q6J("placeholder",e.lcZ(12,23,"NameExample"))("ngModel",o.UserToEdit.Name),e.xp6(4),e.Oqu(e.lcZ(16,25,"Email")),e.xp6(2),e.Q6J("ngModel",o.UserToEdit.Email),e.xp6(3),e.Oqu(e.lcZ(21,27,"Phone")),e.xp6(2),e.Q6J("ngModel",o.UserToEdit.Phone),e.xp6(3),e.Oqu(e.lcZ(26,29,"Language")),e.xp6(2),e.Q6J("ngbTooltip",e.lcZ(29,31,"LangSelector"))("ngModel",o.UserToEdit.Lang),e.xp6(3),e.Q6J("ngForOf",o.translate.getLangs()),e.xp6(2),e.Oqu(e.lcZ(33,33,"Password")),e.xp6(4),e.Q6J("ngModel",o.changepassword),e.xp6(1),e.Q6J("placeholder",e.lcZ(38,35,"EnterNewPassword"))("disabled",!o.changepassword),e.xp6(2),e.Q6J("disabled",t.invalid),e.xp6(1),e.Oqu(e.lcZ(41,37,"Save")),e.xp6(3),e.Oqu(e.lcZ(44,39,"Cancel"))}}function F(n,i){if(1&n){const t=e.EpF();e.TgZ(0,"div",10)(1,"h2"),e._uU(2),e.ALo(3,"translate"),e.qZA(),e.TgZ(4,"button",30),e.NdJ("click",function(){return e.CHM(t),e.oxw().OnUnlinkTwoFactor()}),e._uU(5),e.ALo(6,"translate"),e.qZA()()}2&n&&(e.xp6(2),e.Oqu(e.lcZ(3,2,"SecondFactorEnabled")),e.xp6(3),e.Oqu(e.lcZ(6,4,"Disable")))}function y(n,i){if(1&n){const t=e.EpF();e.TgZ(0,"form",8,31),e.NdJ("ngSubmit",function(){e.CHM(t);const r=e.MAs(1);return e.oxw().OnLinkTwoFactor(r)}),e.TgZ(2,"div",10)(3,"div",11)(4,"h2"),e._uU(5),e.ALo(6,"translate"),e.qZA(),e.TgZ(7,"h4"),e._uU(8),e.ALo(9,"translate"),e.TgZ(10,"a",32),e._uU(11,"Authenticator"),e.qZA(),e._uU(12,": "),e.qZA(),e._UZ(13,"img",33),e.ALo(14,"async"),e.ALo(15,"secureImagePipe"),e.ALo(16,"translate"),e.TgZ(17,"h4"),e._uU(18),e.ALo(19,"translate"),e.qZA(),e.TgZ(20,"label",22),e._uU(21),e.ALo(22,"translate"),e.qZA(),e.TgZ(23,"div",23)(24,"div",24)(25,"input",34),e.NdJ("ngModelChange",function(r){return e.CHM(t),e.oxw().SetTwoFactor=r}),e.qZA()(),e._UZ(26,"input",35),e.ALo(27,"translate"),e.ALo(28,"translate"),e.qZA(),e.TgZ(29,"button",36),e._uU(30),e.ALo(31,"translate"),e.qZA()()()()}if(2&n){const t=e.oxw();e.xp6(5),e.Oqu(e.lcZ(6,12,"SetSecondFactor")),e.xp6(3),e.hij("1. ",e.lcZ(9,14,"ScanWith")," "),e.xp6(2),e.Q6J("href",t.AuthUrl,e.LSH),e.xp6(3),e.Q6J("src",e.lcZ(14,16,e.lcZ(15,18,t.QrUrl)),e.LSH)("alt",e.lcZ(16,20,"QrCode")),e.xp6(5),e.hij("2. ",e.lcZ(19,22,"EnterToken"),""),e.xp6(3),e.Oqu(e.lcZ(22,24,"Token")),e.xp6(4),e.Q6J("ngModel",t.SetTwoFactor),e.xp6(1),e.Q6J("placeholder",e.lcZ(27,26,"TokenFromApp"))("disabled",!t.SetTwoFactor)("ngbTooltip",e.lcZ(28,28,"MinSixChars")),e.xp6(4),e.Oqu(e.lcZ(31,30,"Save"))}}let O=(()=>{class n{constructor(t,o,r,l){this.auth=t,this.datastore=o,this.sitetitle=r,this.translate=l,l.addLangs(c.N.SupportedLangs),l.setDefaultLang(c.N.DefaultLocale)}ngOnDestroy(){this.FetchUser.unsubscribe(),this.RecivedErrorSub.unsubscribe(),this.DataLoading.unsubscribe(),this.SaveUser.unsubscribe(),this.LinkUnlinkTFA.unsubscribe()}ngOnInit(){const t=localStorage.getItem("userLang");this.SwitchLanguage(null!==t?t:c.N.DefaultLocale),this.AuthUrl=c.N.GetAuthenticatorUrl,this.QrUrl=c.N.GetTOTPQRCodeUrl,this.LinkUnlinkTFA=this.datastore.TwoFactorSub.subscribe(o=>{this.SetUserAndTFA(o)}),this.SaveUser=this.datastore.UserUpdateInsert.subscribe(o=>{this.SetUserAndTFA(o)}),this.FetchUser=this.datastore.CurrentUserFetch.subscribe(o=>{this.SetUserAndTFA(o)}),this.RecivedErrorSub=this.datastore.RecivedError.subscribe(o=>{this.ShowMessage=!0,this.ResponseFromBackend=o,setTimeout(()=>{this.ShowMessage=!1,(401===o.Error.Code||403===o.Error.Code||407===o.Error.Code)&&this.auth.SignOut()},c.N.MessageTimeout),o&&(this.MessageType=200===o.Error.Code?"success":"danger")}),this.DataLoading=this.datastore.LoadingData.subscribe(o=>{this.IsLoading=o}),this.datastore.FetchCurrentUser()}SetUserAndTFA(t){this.UserToEdit=t,this.UserToEdit&&(this.TwoFactorEnabled=this.UserToEdit.SecondFactor)}OnSaveClick(t){if(t.valid){if(t.value.changepassword&&0===t.value.newpassword.length)return;this.UserToEdit.Email=t.value.useremail,this.UserToEdit.Name=t.value.username,this.UserToEdit.Phone=t.value.userphone,this.UserToEdit.Lang=t.value.userlanguage,this.datastore.SaveCurrentUser(this.UserToEdit,t.value.changepassword,t.value.newpassword)}}OnLinkTwoFactor(t){t.valid&&this.datastore.LinkTwoFactor(t.value.passkey,this.UserToEdit)}OnUnlinkTwoFactor(){this.UserToEdit.SecondFactor&&this.datastore.UnlinkTwoFactor(this.UserToEdit)}SwitchLanguage(t){this.translate.use(t),localStorage.setItem("userLang",t),this.auth.ChangeLocale(t),this.translate.get("WebsiteTitleText",t).subscribe({next:o=>{this.sitetitle.setTitle(o)},error:o=>{console.log(o)}})}}return n.\u0275fac=function(t){return new(t||n)(e.Y36(T.e),e.Y36(b.Z),e.Y36(Z.Dx),e.Y36(p.sK))},n.\u0275cmp=e.Xpm({type:n,selectors:[["app-user-profile"]],decls:6,vars:5,consts:[[1,"text-center"],["class","spinner-border","role","status",4,"ngIf"],[3,"type","close",4,"ngIf"],[3,"ngSubmit",4,"ngIf"],["class","form-group",4,"ngIf"],["role","status",1,"spinner-border"],[1,"visually-hidden-focusable"],[3,"type","close"],[3,"ngSubmit"],["UserProfileForm","ngForm"],[1,"form-group"],[2,"margin","3px"],["for","name"],["type","text","id","name","name","username","required","",1,"form-control","mb-1",3,"placeholder","ngModel"],["for","email"],["type","email","id","email","placeholder","exampe@example.com","name","useremail","required","","email","",1,"form-control","mb-1",3,"ngModel"],["for","phone"],["type","tel","id","phone","placeholder","+7 (965) 777-77-77","name","userphone","pattern","^((8|\\+7)[\\- ]?)?(\\(?\\d{3,4}\\)?[\\- ]?)?[\\d\\- ]{5,10}$",1,"form-control","mb-1",3,"ngModel"],["for","lang"],["id","lang","name","userlanguage","required","",1,"form-select","mb-1",3,"ngbTooltip","ngModel","change"],["selectedLang",""],[3,"value","selected",4,"ngFor","ngForOf"],["for","newpassword"],[1,"input-group","mb-3"],[1,"input-group-text"],["type","checkbox","id","changepassword","name","changepassword",1,"form-check-input","mt-0",3,"ngModel","ngModelChange"],["type","password","id","newpassword","name","newpassword","ngModel","",1,"form-control",3,"placeholder","disabled"],["type","submit",1,"btn","btn-outline-primary",3,"disabled"],["type","button","routerLink","..",1,"btn","btn-outline-danger",2,"margin-left","3px"],[3,"value","selected"],["type","button",1,"btn","btn-outline-danger",3,"click"],["UserTwoFactorForm","ngForm"],[3,"href"],[3,"src","alt"],["type","checkbox","id","enabletwofactor","name","enabletwofactor","required","",3,"ngModel","ngModelChange"],["type","text","name","passkey","inputmode","numeric","pattern","[0-9]*","autocomplete","one-time-code","id","newpasskey","ngModel","","required","","minlength","6",1,"form-control",3,"placeholder","disabled","ngbTooltip"],["type","submit",1,"btn","btn-outline-primary"]],template:function(t,o){1&t&&(e.TgZ(0,"div",0),e.YNc(1,C,4,3,"div",1),e.qZA(),e.YNc(2,M,3,4,"ngb-alert",2),e.YNc(3,E,45,41,"form",3),e.YNc(4,F,7,6,"div",4),e.YNc(5,y,32,32,"form",3)),2&t&&(e.xp6(1),e.Q6J("ngIf",o.IsLoading),e.xp6(1),e.Q6J("ngIf",o.ShowMessage),e.xp6(1),e.Q6J("ngIf",o.UserToEdit),e.xp6(1),e.Q6J("ngIf",o.UserToEdit&&o.TwoFactorEnabled),e.xp6(1),e.Q6J("ngIf",o.UserToEdit&&!o.TwoFactorEnabled))},directives:[u.O5,g.xm,s._Y,s.JL,s.F,s.Fj,s.Q7,s.JJ,s.On,s.on,s.c5,s.EJ,g._L,u.sg,s.YN,s.Kr,s.Wl,U.rH,s.Zs,s.wO],pipes:[p.X$,v.J,u.Ov,L],styles:["input.ng-invalid.ng-touched[_ngcontent-%COMP%]{border-color:#ff0a1d!important;box-shadow:none!important}input.ng-invalid.ng-touched[_ngcontent-%COMP%]:focus{box-shadow:0 0 0 .2rem #ff0a1d40!important}input.ng-valid.ng-touched[_ngcontent-%COMP%]{border-color:#67c07b!important;box-shadow:none!important}input.ng-valid.ng-touched[_ngcontent-%COMP%]:focus{box-shadow:0 0 0 .2rem #34b30d40!important}select.ng-invalid.ng-touched[_ngcontent-%COMP%]{border-color:#ff0a1d!important;box-shadow:none!important}select.ng-invalid.ng-touched[_ngcontent-%COMP%]:focus{box-shadow:0 0 0 .2rem #ff0a1d40!important}select.ng-valid.ng-touched[_ngcontent-%COMP%]{border-color:#67c07b!important;box-shadow:none!important}select.ng-valid.ng-touched[_ngcontent-%COMP%]:focus{box-shadow:0 0 0 .2rem #34b30d40!important}.image-placeholder[_ngcontent-%COMP%]{background-color:#eee;display:flex;height:333px;margin:5px;border-radius:3px;max-width:592px}.image-placeholder[_ngcontent-%COMP%] > h4[_ngcontent-%COMP%]{align-self:center;text-align:center;width:100%}"]}),n})();var S=a(319),J=a(4466);const q=[{path:"",component:O}];let N=(()=>{class n{}return n.\u0275fac=function(t){return new(t||n)},n.\u0275mod=e.oAB({type:n}),n.\u0275inj=e.cJS({imports:[[J.m,u.ez,p.aw.forRoot({loader:{provide:p.Zw,useFactory:x,deps:[h.eN]}}),s.u5,g._A,U.Bz.forChild(q),g.HK]]}),n})();function x(n){return new S.w(n)}}}]);