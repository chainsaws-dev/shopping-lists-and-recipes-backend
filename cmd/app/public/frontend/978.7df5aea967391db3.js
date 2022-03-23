"use strict";(self.webpackChunkshopping_lists_and_recipes=self.webpackChunkshopping_lists_and_recipes||[]).push([[978],{3978:(v,p,i)=>{i.r(p),i.d(p,{AuthFeatureModule:()=>O,httpTranslateLoader:()=>f});var d=i(1048),m=i(3481),g=i(2340),t=i(9724),h=i(384),l=i(1404),r=i(4190),o=i(1659),u=i(2719);function _(n,a){1&n&&(t.TgZ(0,"div",9),t._UZ(1,"input",25),t.ALo(2,"translate"),t.qZA()),2&n&&(t.xp6(1),t.Q6J("placeholder",t.lcZ(2,1,"Name")))}function x(n,a){if(1&n&&(t.TgZ(0,"option",26),t._uU(1),t.qZA()),2&n){const e=a.$implicit,s=t.oxw(2);t.Q6J("value",e)("selected",e===s.translate.currentLang),t.xp6(1),t.hij(" ",e," ")}}function A(n,a){1&n&&(t.TgZ(0,"a",27),t._uU(1),t.ALo(2,"translate"),t.qZA()),2&n&&(t.xp6(1),t.Oqu(t.lcZ(2,1,"ForgotYourPassword")))}function L(n,a){1&n&&(t.TgZ(0,"a",28),t._uU(1),t.ALo(2,"translate"),t.qZA()),2&n&&(t.xp6(1),t.Oqu(t.lcZ(2,1,"NoConfirmationEmail")))}function Z(n,a){if(1&n){const e=t.EpF();t.TgZ(0,"ngb-alert",31),t.NdJ("close",function(){return t.CHM(e),t.oxw(3).ShowMessage=!1}),t._uU(1),t.qZA()}if(2&n){const e=t.oxw(3);t.Q6J("type",e.MessageType),t.xp6(1),t.Oqu(e.ResponseFromBackend.Error.Message)}}function b(n,a){if(1&n&&(t.TgZ(0,"div",29),t.YNc(1,Z,2,2,"ngb-alert",30),t.qZA()),2&n){const e=t.oxw(2);t.xp6(1),t.Q6J("ngIf",e.ShowMessage)}}function T(n,a){if(1&n){const e=t.EpF();t.TgZ(0,"div",4)(1,"div",5)(2,"h3"),t._uU(3),t.ALo(4,"translate"),t.qZA()(),t.TgZ(5,"div",6)(6,"form",7,8),t.NdJ("ngSubmit",function(){t.CHM(e);const c=t.MAs(7);return t.oxw().OnSubmitForm(c)}),t.TgZ(8,"div",9),t._UZ(9,"input",10),t.ALo(10,"translate"),t.ALo(11,"translate"),t.qZA(),t.YNc(12,_,3,3,"div",11),t.TgZ(13,"div",9),t._UZ(14,"input",12),t.ALo(15,"translate"),t.ALo(16,"translate"),t.qZA(),t.TgZ(17,"div",13)(18,"div",14)(19,"button",15),t._uU(20),t.ALo(21,"translate"),t.qZA(),t.TgZ(22,"button",16),t.NdJ("click",function(){return t.CHM(e),t.oxw().onSwitchMode()}),t._uU(23),t.ALo(24,"translate"),t.qZA(),t.TgZ(25,"select",17,18),t.NdJ("change",function(){t.CHM(e);const c=t.MAs(26);return t.oxw().SwitchLanguage(c.value)}),t.ALo(27,"translate"),t.YNc(28,x,2,3,"option",19),t.qZA()()()()(),t.TgZ(29,"div",20)(30,"div",21),t.YNc(31,A,3,3,"a",22),t.YNc(32,L,3,3,"a",23),t.qZA()(),t.YNc(33,b,2,1,"div",24),t.qZA()}if(2&n){const e=t.oxw();t.xp6(3),t.Oqu(t.lcZ(4,13,e.LoginMode?"SignIn":"SignUp")),t.xp6(6),t.Q6J("placeholder",t.lcZ(10,15,"Email"))("ngbTooltip",t.lcZ(11,17,"RealEmailRequired")),t.xp6(3),t.Q6J("ngIf",!e.LoginMode),t.xp6(2),t.Q6J("placeholder",t.lcZ(15,19,"Password"))("ngbTooltip",t.lcZ(16,21,"MinSixChars")),t.xp6(6),t.Oqu(t.lcZ(21,23,e.LoginMode?"SignIn":"SignUp")),t.xp6(3),t.hij(" ",t.lcZ(24,25,e.LoginMode?"SignUp":"SignIn"),""),t.xp6(2),t.Q6J("ngbTooltip",t.lcZ(27,27,"LangSelector")),t.xp6(3),t.Q6J("ngForOf",e.translate.getLangs()),t.xp6(3),t.Q6J("ngIf",e.LoginMode),t.xp6(1),t.Q6J("ngIf",!e.LoginMode),t.xp6(1),t.Q6J("ngIf",e.ShowMessage)}}function M(n,a){1&n&&(t.TgZ(0,"div",32)(1,"span",33),t._uU(2),t.ALo(3,"translate"),t.qZA()()),2&n&&(t.xp6(2),t.Oqu(t.lcZ(3,1,"Loading")))}let C=(()=>{class n{constructor(e,s,c){this.authservice=e,this.router=s,this.translate=c,this.LoginMode=!0,this.IsLoading=!1,c.addLangs(g.N.SupportedLangs),c.setDefaultLang(g.N.DefaultLocale)}ngOnInit(){const e=localStorage.getItem("userLang");this.SwitchLanguage(null!==e?e:g.N.DefaultLocale),this.authservice.CheckRegistered()&&this.Redirect(),this.authErrSub=this.authservice.AuthErrorSub.subscribe(s=>{this.ShowMessage=!0,this.ResponseFromBackend=s,setTimeout(()=>this.ShowMessage=!1,5e3),s&&(this.MessageType=200===s.Error.Code?"success":"danger"),this.IsLoading=!1}),this.loginResultSub=this.authservice.AuthResultSub.subscribe(s=>{this.loggedIn=s,this.IsLoading=!1,s&&(this.authservice.HaveToCheckSecondFactor()?this.GoToSecondFactor():this.Redirect())})}SwitchLanguage(e){this.translate.use(e),localStorage.setItem("userLang",e)}ngOnDestroy(){this.loginResultSub.unsubscribe(),this.authErrSub.unsubscribe()}onSwitchMode(){this.LoginMode=!this.LoginMode}OnSubmitForm(e){e.invalid?(this.ShowMessage=!0,this.ResponseFromBackend=new m.iQ(400,this.translate.instant("IncorrectDataInput")),this.MessageType="danger",setTimeout(()=>this.ShowMessage=!1,5e3)):(this.IsLoading=!0,this.LoginMode?this.authservice.SignIn(e.value.email,e.value.password):this.authservice.SignUp(e.value.email,e.value.name,e.value.password),e.reset())}Redirect(){this.router.navigate(["/recipes"])}GoToSecondFactor(){this.router.navigate(["/totp"])}}return n.\u0275fac=function(e){return new(e||n)(t.Y36(h.e),t.Y36(l.F0),t.Y36(r.sK))},n.\u0275cmp=t.Xpm({type:n,selectors:[["app-auth"]],decls:4,vars:2,consts:[[1,"main-parent-login"],["class","card mx-auto","style","margin-top: 3px;",4,"ngIf"],[1,"text-center"],["class","spinner-border","role","status",4,"ngIf"],[1,"card","mx-auto",2,"margin-top","3px"],[1,"card-header"],[1,"card-body"],[3,"ngSubmit"],["SignupForm","ngForm"],[1,"input-group","mb-3"],["type","email","name","email","placement","right","ngModel","","email","","required","",1,"form-control",3,"placeholder","ngbTooltip"],["class","input-group mb-3",4,"ngIf"],["type","password","name","password","ngModel","","required","","minlength","6","placement","right",1,"form-control",3,"placeholder","ngbTooltip"],[1,"float-right"],[1,"input-group"],["type","submit",1,"btn","btn-outline-primary","me-1"],["type","button",1,"btn","btn-outline-primary",3,"click"],[1,"form-select",3,"ngbTooltip","change"],["selectedLang",""],[3,"value","selected",4,"ngFor","ngForOf"],[1,"card-footer"],[1,"d-flex","justify-content-center"],["href","/reset-password",4,"ngIf"],["href","/confirm-email",4,"ngIf"],["style","padding: 3px",4,"ngIf"],["type","text","name","name","ngModel","",1,"form-control",3,"placeholder"],[3,"value","selected"],["href","/reset-password"],["href","/confirm-email"],[2,"padding","3px"],[3,"type","close",4,"ngIf"],[3,"type","close"],["role","status",1,"spinner-border"],[1,"visually-hidden-focusable"]],template:function(e,s){1&e&&(t.TgZ(0,"div",0),t.YNc(1,T,34,29,"div",1),t.TgZ(2,"div",2),t.YNc(3,M,4,3,"div",3),t.qZA()()),2&e&&(t.xp6(1),t.Q6J("ngIf",!s.IsLoading),t.xp6(2),t.Q6J("ngIf",s.IsLoading))},directives:[d.O5,o._Y,o.JL,o.F,o.Fj,o.JJ,o.On,o.on,o.Q7,u._L,o.wO,d.sg,o.YN,o.Kr,u.xm],pipes:[r.X$],styles:[".main-parent-login[_ngcontent-%COMP%]{position:absolute;top:50%;left:50%;transform:translate(-50%,-50%)}.card[_ngcontent-%COMP%]{width:300px}.remember[_ngcontent-%COMP%]   input[_ngcontent-%COMP%]{width:20px;height:20px;margin-left:15px;margin-right:5px}input.ng-invalid.ng-touched[_ngcontent-%COMP%]{border-color:#ff0a1d!important;box-shadow:none!important}input.ng-invalid.ng-touched[_ngcontent-%COMP%]:focus{box-shadow:0 0 0 .2rem #ff0a1d40!important}input.ng-valid.ng-touched[_ngcontent-%COMP%]{border-color:#67c07b!important;box-shadow:none!important}input.ng-valid.ng-touched[_ngcontent-%COMP%]:focus{box-shadow:0 0 0 .2rem #34b30d40!important}"]}),n})();var S=i(4466),w=i(9725),I=i(319);const y=[{path:"",component:C}];let O=(()=>{class n{}return n.\u0275fac=function(e){return new(e||n)},n.\u0275mod=t.oAB({type:n}),n.\u0275inj=t.cJS({imports:[[d.ez,o.u5,l.Bz.forChild(y),r.aw.forRoot({loader:{provide:r.Zw,useFactory:f,deps:[w.eN]}}),S.m,u._A,u.HK]]}),n})();function f(n){return new I.w(n)}},3481:(v,p,i)=>{i.d(p,{iQ:()=>g,o7:()=>d,tl:()=>h});class d{constructor(r,o){this.Name=r,this.Amount=o}}class g{constructor(r,o){this.Error=new t(r,o)}}class t{constructor(r,o){this.Code=r,this.Message=o}}class h{constructor(r,o,u){this.Total=r,this.Limit=o,this.Offset=u}}}}]);