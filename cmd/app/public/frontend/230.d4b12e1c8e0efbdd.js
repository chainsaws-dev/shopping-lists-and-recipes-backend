"use strict";(self.webpackChunkshopping_lists_and_recipes=self.webpackChunkshopping_lists_and_recipes||[]).push([[230],{2230:(y,l,s)=>{s.r(l),s.d(l,{TotpModule:()=>Z,httpTranslateLoader:()=>h});var u=s(1048),g=s(1404),m=s(4466),a=s(1659),c=s(2719),d=s(2340),t=s(9724),f=s(384),p=s(4190);function v(n,i){if(1&n){const e=t.EpF();t.TgZ(0,"ngb-alert",17),t.NdJ("close",function(){return t.CHM(e),t.oxw(3).ShowMessage=!1}),t._uU(1),t.qZA()}if(2&n){const e=t.oxw(3);t.Q6J("type",e.MessageType),t.xp6(1),t.Oqu(e.ResponseFromBackend.Error.Message)}}function x(n,i){if(1&n&&(t.TgZ(0,"div",15),t.YNc(1,v,2,2,"ngb-alert",16),t.qZA()),2&n){const e=t.oxw(2);t.xp6(1),t.Q6J("ngIf",e.ShowMessage)}}function b(n,i){if(1&n){const e=t.EpF();t.TgZ(0,"div",4)(1,"div",5)(2,"h3"),t._uU(3),t.ALo(4,"translate"),t.qZA()(),t.TgZ(5,"div",6)(6,"form",7,8),t.NdJ("ngSubmit",function(){t.CHM(e);const r=t.MAs(7);return t.oxw().OnSubmitForm(r)}),t.TgZ(8,"div",9),t._UZ(9,"input",10),t.ALo(10,"translate"),t.ALo(11,"translate"),t.qZA(),t.TgZ(12,"div",11)(13,"div",12)(14,"button",13),t._uU(15),t.ALo(16,"translate"),t.qZA()()()()(),t.YNc(17,x,2,1,"div",14),t.qZA()}if(2&n){const e=t.MAs(7),o=t.oxw();t.xp6(3),t.Oqu(t.lcZ(4,6,"SecondFactor")),t.xp6(6),t.Q6J("placeholder",t.lcZ(10,8,"Passkey"))("ngbTooltip",t.lcZ(11,10,"MinSixChars")),t.xp6(5),t.Q6J("disabled",e.invalid),t.xp6(1),t.Oqu(t.lcZ(16,12,"Check")),t.xp6(2),t.Q6J("ngIf",o.ShowMessage)}}function T(n,i){1&n&&(t.TgZ(0,"div",18)(1,"span",19),t._uU(2),t.ALo(3,"translate"),t.qZA()()),2&n&&(t.xp6(2),t.Oqu(t.lcZ(3,1,"Loading")))}let S=(()=>{class n{constructor(e,o,r){this.authservice=e,this.router=o,this.translate=r,this.LoginMode=!0,this.IsLoading=!1,r.addLangs(d.N.SupportedLangs),r.setDefaultLang(d.N.DefaultLocale)}ngOnDestroy(){this.SfResultSub.unsubscribe(),this.SfErrSub.unsubscribe()}ngOnInit(){this.authservice.CheckRegistered()&&this.Redirect();const e=localStorage.getItem("userLang");this.SwitchLanguage(null!==e?e:d.N.DefaultLocale),this.SfErrSub=this.authservice.SfErrorSub.subscribe(o=>{this.ShowMessage=!0,this.ResponseFromBackend=o,setTimeout(()=>this.ShowMessage=!1,5e3),o&&(this.MessageType=200===o.Error.Code?"success":"danger"),this.IsLoading=!1}),this.SfResultSub=this.authservice.SfResultSub.subscribe(o=>{this.IsLoading=!1,o&&this.Redirect()})}OnSubmitForm(e){this.IsLoading=!0,this.authservice.SecondFactorCheck(e.value.passkey),e.reset()}Redirect(){this.router.navigate(["/recipes"])}SwitchLanguage(e){this.translate.use(e),localStorage.setItem("userLang",e)}}return n.\u0275fac=function(e){return new(e||n)(t.Y36(f.e),t.Y36(g.F0),t.Y36(p.sK))},n.\u0275cmp=t.Xpm({type:n,selectors:[["app-totp"]],decls:4,vars:2,consts:[[1,"main-parent-login"],["class","card mx-auto","style","margin-top: 3px;",4,"ngIf"],[1,"text-center"],["class","spinner-border","role","status",4,"ngIf"],[1,"card","mx-auto",2,"margin-top","3px"],[1,"card-header"],[1,"card-body"],[3,"ngSubmit"],["SignupForm","ngForm"],[1,"input-group","form-group"],["type","text","name","passkey","inputmode","numeric","pattern","[0-9]*","autocomplete","one-time-code","ngModel","","required","","minlength","6","placement","right",1,"form-control","mb-3",3,"placeholder","ngbTooltip"],[1,"form-group","float-right"],[1,"input-group"],["type","submit",1,"btn","btn-outline-primary",3,"disabled"],["style","padding: 3px",4,"ngIf"],[2,"padding","3px"],[3,"type","close",4,"ngIf"],[3,"type","close"],["role","status",1,"spinner-border"],[1,"visually-hidden-focusable"]],template:function(e,o){1&e&&(t.TgZ(0,"div",0),t.YNc(1,b,18,14,"div",1),t.TgZ(2,"div",2),t.YNc(3,T,4,3,"div",3),t.qZA()()),2&e&&(t.xp6(1),t.Q6J("ngIf",!o.IsLoading),t.xp6(2),t.Q6J("ngIf",o.IsLoading))},directives:[u.O5,a._Y,a.JL,a.F,a.Fj,a.c5,a.JJ,a.On,a.Q7,a.wO,c._L,c.xm],pipes:[p.X$],styles:[".main-parent-login[_ngcontent-%COMP%]{position:absolute;top:50%;left:50%;transform:translate(-50%,-50%)}.card[_ngcontent-%COMP%]{width:300px}.remember[_ngcontent-%COMP%]   input[_ngcontent-%COMP%]{width:20px;height:20px;margin-left:15px;margin-right:5px}input.ng-invalid.ng-touched[_ngcontent-%COMP%]{border-color:#ff0a1d!important;box-shadow:none!important}input.ng-invalid.ng-touched[_ngcontent-%COMP%]:focus{box-shadow:0 0 0 .2rem #ff0a1d40!important}input.ng-valid.ng-touched[_ngcontent-%COMP%]{border-color:#67c07b!important;box-shadow:none!important}input.ng-valid.ng-touched[_ngcontent-%COMP%]:focus{box-shadow:0 0 0 .2rem #34b30d40!important}"]}),n})();var C=s(319),M=s(9725);const L=[{path:"",component:S}];let Z=(()=>{class n{}return n.\u0275fac=function(e){return new(e||n)},n.\u0275mod=t.oAB({type:n}),n.\u0275inj=t.cJS({imports:[[u.ez,a.u5,g.Bz.forChild(L),p.aw.forRoot({loader:{provide:p.Zw,useFactory:h,deps:[M.eN]}}),m.m,c._A,c.HK]]}),n})();function h(n){return new C.w(n)}}}]);