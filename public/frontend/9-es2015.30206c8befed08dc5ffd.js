(window.webpackJsonp=window.webpackJsonp||[]).push([[9],{jkDv:function(e,t,n){"use strict";n.r(t),n.d(t,"AdminModule",(function(){return Q}));var i=n("ofXK"),s=n("3Pt+"),o=n("1kSV"),r=n("tyNb"),a=n("P+IX"),c=n("dxYa"),b=n("AytR"),d=n("lJxs"),l=n("fXoL"),u=n("GXvH");let p=(()=>{class e{constructor(e){this.datastorageservice=e}resolve(e,t){const n=e.params.id;return this.datastorageservice.FetchUsersList(e.params.pn,b.a.AdminUserListPageSize).pipe(Object(d.a)(e=>e.Users[n]))}}return e.\u0275fac=function(t){return new(t||e)(l.Xb(u.a))},e.\u0275prov=l.Gb({token:e,factory:e.\u0275fac,providedIn:"root"}),e})(),g=(()=>{class e{constructor(){}ngOnInit(){}}return e.\u0275fac=function(t){return new(t||e)},e.\u0275cmp=l.Eb({type:e,selectors:[["app-admin"]],decls:11,vars:0,consts:[[1,"nav","nav-tabs"],[1,"active"],["routerLink","/admin/users","routerLinkActive","active",1,"nav-link"],["routerLink","/admin/sessions","routerLinkActive","active",1,"nav-link"],["routerLink","/admin/media","routerLinkActive","active",1,"nav-link"]],template:function(e,t){1&e&&(l.Pb(0,"ul",0),l.Pb(1,"li",1),l.Pb(2,"a",2),l.xc(3,"Users"),l.Ob(),l.Ob(),l.Pb(4,"li"),l.Pb(5,"a",3),l.xc(6,"Sessions"),l.Ob(),l.Ob(),l.Pb(7,"li"),l.Pb(8,"a",4),l.xc(9,"Media"),l.Ob(),l.Ob(),l.Ob(),l.Lb(10,"router-outlet"))},directives:[r.f,r.e,r.h],styles:[""]}),e})();var h=n("ggGY"),m=n("c7dm");function f(e,t){1&e&&(l.Pb(0,"div",4),l.Pb(1,"span",5),l.xc(2,"Loading..."),l.Ob(),l.Ob())}function v(e,t){if(1&e){const e=l.Qb();l.Pb(0,"ngb-alert",6),l.ac("close",(function(){return l.oc(e),l.cc().ShowMessage=!1})),l.xc(1),l.Ob()}if(2&e){const e=l.cc();l.fc("type",e.MessageType),l.xb(1),l.zc(" ",e.ResponseFromBackend.Error.Message,"")}}function x(e,t){1&e&&(l.Pb(0,"h2"),l.xc(1,"Edit"),l.Ob())}function P(e,t){1&e&&(l.Pb(0,"h2"),l.xc(1,"Create"),l.Ob())}function O(e,t){if(1&e){const e=l.Qb();l.Pb(0,"form",7,8),l.ac("ngSubmit",(function(){l.oc(e);const t=l.nc(1);return l.cc().OnSaveClick(t)})),l.Pb(2,"div",9),l.Pb(3,"div",10),l.vc(4,x,2,0,"h2",11),l.vc(5,P,2,0,"h2",11),l.Pb(6,"div",9),l.Pb(7,"label",12),l.xc(8,"Name"),l.Ob(),l.Lb(9,"input",13),l.Ob(),l.Pb(10,"div",9),l.Pb(11,"label",14),l.xc(12,"Email"),l.Ob(),l.Lb(13,"input",15),l.Ob(),l.Pb(14,"div",9),l.Pb(15,"label",16),l.xc(16,"Phone"),l.Ob(),l.Lb(17,"input",17),l.Ob(),l.Pb(18,"div",9),l.Pb(19,"label",18),l.xc(20,"Select role:"),l.Ob(),l.Pb(21,"select",19),l.Pb(22,"option",20),l.xc(23,"Guest"),l.Ob(),l.Pb(24,"option",21),l.xc(25,"Admin"),l.Ob(),l.Ob(),l.Ob(),l.Pb(26,"label",22),l.xc(27,"Password"),l.Ob(),l.Pb(28,"div",23),l.Pb(29,"div",24),l.Pb(30,"div",25),l.Pb(31,"input",26),l.ac("ngModelChange",(function(t){return l.oc(e),l.cc().changepassword=t})),l.Ob(),l.Ob(),l.Ob(),l.Lb(32,"input",27),l.Ob(),l.Pb(33,"button",28),l.xc(34,"Save"),l.Ob(),l.Pb(35,"button",29),l.xc(36,"Cancel"),l.Ob(),l.Ob(),l.Ob(),l.Ob()}if(2&e){const e=l.nc(1),t=l.cc();l.xb(4),l.fc("ngIf",t.editmode),l.xb(1),l.fc("ngIf",!t.editmode),l.xb(4),l.fc("ngModel",t.UserToEdit.Name),l.xb(4),l.fc("ngModel",t.UserToEdit.Email),l.xb(4),l.fc("ngModel",t.UserToEdit.Phone),l.xb(4),l.fc("ngModel",t.UserToEdit.Role),l.xb(10),l.fc("ngModel",t.changepassword),l.xb(1),l.fc("disabled",!t.changepassword),l.xb(1),l.fc("disabled",e.invalid)}}let S=(()=>{class e{constructor(e,t,n,i){this.AdminServ=e,this.activatedroute=t,this.router=n,this.datastore=i}ngOnDestroy(){this.DataLoading.unsubscribe(),this.RecivedErrorSub.unsubscribe(),this.DatabaseUpdated&&this.DatabaseUpdated.unsubscribe()}ngOnInit(){this.activatedroute.params.subscribe(e=>{this.editmode=null!=e.id,this.editmode?(this.index=+e.id,this.UserToEdit=this.AdminServ.GetUserById(this.index)):(this.changepassword=!0,this.UserToEdit=new h.a("guest_role_read_only","","","")),this.AdminServ.CurrentSelectedItem=this.UserToEdit}),this.RecivedErrorSub=this.datastore.RecivedError.subscribe(e=>{switch(this.ShowMessage=!0,this.ResponseFromBackend=e,setTimeout(()=>this.ShowMessage=!1,5e3),e.Error.Code){case 200:this.MessageType="success";break;default:this.MessageType="danger"}}),this.DataLoading=this.datastore.LoadingData.subscribe(e=>{this.IsLoading=e})}OnSaveClick(e){if(e.valid){if(e.value.changepassword&&0===e.value.newpassword.length)return;this.UserToEdit.Email=e.value.useremail,this.UserToEdit.Name=e.value.username,this.UserToEdit.Phone=e.value.userphone,this.UserToEdit.Role=e.value.roles,this.datastore.SaveUser(this.UserToEdit,e.value.changepassword,e.value.newpassword),this.DatabaseUpdated=this.datastore.UserUpdateInsert.subscribe(e=>{this.editmode?(this.UserToEdit=e,this.AdminServ.UpdateExistingUser(this.UserToEdit,this.index)):(this.UserToEdit=e,this.AdminServ.AddNewUser(this.UserToEdit)),setTimeout(()=>this.router.navigate(["../"],{relativeTo:this.activatedroute,queryParamsHandling:"merge"}),1e3)})}}}return e.\u0275fac=function(t){return new(t||e)(l.Kb(m.a),l.Kb(r.a),l.Kb(r.c),l.Kb(u.a))},e.\u0275cmp=l.Eb({type:e,selectors:[["app-user-edit"]],decls:4,vars:3,consts:[[1,"text-center"],["class","spinner-border","role","status",4,"ngIf"],[3,"type","close",4,"ngIf"],[3,"ngSubmit",4,"ngIf"],["role","status",1,"spinner-border"],[1,"sr-only"],[3,"type","close"],[3,"ngSubmit"],["UserEditForm","ngForm"],[1,"form-group"],[2,"margin","3px"],[4,"ngIf"],["for","name"],["type","text","id","name","placeholder","example","name","username","required","",1,"form-control","mb-1",3,"ngModel"],["for","email"],["type","email","id","email","placeholder","exampe@example.com","name","useremail","required","","email","",1,"form-control","mb-1",3,"ngModel"],["for","phone"],["type","tel","id","phone","placeholder","+7 (965) 777-77-77","name","userphone","pattern","^((8|\\+7)[\\- ]?)?(\\(?\\d{3,4}\\)?[\\- ]?)?[\\d\\- ]{5,10}$",1,"form-control","mb-1",3,"ngModel"],["for","roles"],["name","roles","id","roles","required","",1,"form-control",3,"ngModel"],["value","guest_role_read_only"],["value","admin_role_CRUD"],["for","newpassword"],[1,"input-group","mb-3"],[1,"input-group-prepend"],[1,"input-group-text"],["type","checkbox","id","changepassword","name","changepassword",3,"ngModel","ngModelChange"],["type","password","id","newpassword","name","newpassword","placeholder","Enter new password","ngModel","",1,"form-control",3,"disabled"],["type","submit",1,"btn","btn-outline-primary",3,"disabled"],["type","button","routerLink","..",1,"btn","btn-outline-danger",2,"margin-left","3px"]],template:function(e,t){1&e&&(l.Pb(0,"div",0),l.vc(1,f,3,0,"div",1),l.Ob(),l.vc(2,v,2,2,"ngb-alert",2),l.vc(3,O,37,9,"form",3)),2&e&&(l.xb(1),l.fc("ngIf",t.IsLoading),l.xb(1),l.fc("ngIf",t.ShowMessage),l.xb(1),l.fc("ngIf",t.UserToEdit))},directives:[i.k,o.a,s.r,s.i,s.j,s.b,s.o,s.h,s.k,s.c,s.n,s.p,s.l,s.q,s.a,r.d],styles:["input.ng-invalid.ng-touched[_ngcontent-%COMP%]{border-color:#ff0a1d!important;box-shadow:none!important}input.ng-invalid.ng-touched[_ngcontent-%COMP%]:focus{box-shadow:0 0 0 .2rem rgba(255,10,29,.25)!important}input.ng-valid.ng-touched[_ngcontent-%COMP%]{border-color:#67c07b!important;box-shadow:none!important}input.ng-valid.ng-touched[_ngcontent-%COMP%]:focus{box-shadow:0 0 0 .2rem rgba(52,179,13,.25)!important}select.ng-invalid.ng-touched[_ngcontent-%COMP%]{border-color:#ff0a1d!important;box-shadow:none!important}select.ng-invalid.ng-touched[_ngcontent-%COMP%]:focus{box-shadow:0 0 0 .2rem rgba(255,10,29,.25)!important}select.ng-valid.ng-touched[_ngcontent-%COMP%]{border-color:#67c07b!important;box-shadow:none!important}select.ng-valid.ng-touched[_ngcontent-%COMP%]:focus{box-shadow:0 0 0 .2rem rgba(52,179,13,.25)!important}.image-placeholder[_ngcontent-%COMP%]{background-color:#eee;display:flex;height:333px;margin:5px;border-radius:3px;max-width:592px}.image-placeholder[_ngcontent-%COMP%] > h4[_ngcontent-%COMP%]{align-self:center;text-align:center;width:100%}"]}),e})();function w(e,t){if(1&e){const e=l.Qb();l.Pb(0,"ngb-alert",12),l.ac("close",(function(){return l.oc(e),l.cc(2).ShowMessage=!1})),l.xc(1),l.Ob()}if(2&e){const e=l.cc(2);l.fc("type",e.MessageType),l.xb(1),l.zc(" ",e.ResponseFromBackend.Error.Message,"")}}function y(e,t){if(1&e){const e=l.Qb();l.Pb(0,"li",13),l.ac("click",(function(){l.oc(e);const n=t.$implicit;return l.cc(2).AdminServ.SelectItemUsersList(n)})),l.Pb(1,"div",14),l.Pb(2,"p",15),l.xc(3),l.Ob(),l.Pb(4,"div",16),l.Pb(5,"button",17),l.xc(6,"Edit"),l.Ob(),l.Pb(7,"button",18),l.ac("click",(function(){l.oc(e);const n=t.$implicit,i=t.index;return l.cc(2).OnDeleteUser(n,i)})),l.xc(8,"Delete"),l.Ob(),l.Ob(),l.Ob(),l.Ob()}if(2&e){const e=t.$implicit,n=t.index,i=l.cc(2);l.fc("ngClass",i.AdminServ.IsCurrentSelected(e)?"active":""),l.xb(3),l.zc(" ",e.Email,""),l.xb(2),l.fc("routerLink",n)}}function M(e,t){if(1&e){const e=l.Qb();l.Pb(0,"ngb-pagination",19),l.ac("pageChange",(function(t){return l.oc(e),l.cc(2).OnPageChanged(t)})),l.Ob()}if(2&e){const e=l.cc(2);l.fc("collectionSize",e.usCollectionSize)("pageSize",e.usPageSize)("page",e.usCurrentPage)("rotate",!0)("boundaryLinks",!0)("maxSize",11)}}const C=function(e){return["/admin/users",e,"new"]};function k(e,t){if(1&e&&(l.Pb(0,"div",3),l.Pb(1,"div",4),l.Pb(2,"div",5),l.Pb(3,"div",6),l.Pb(4,"button",7),l.xc(5,"Add"),l.Ob(),l.Ob(),l.Ob(),l.vc(6,w,2,2,"ngb-alert",8),l.Pb(7,"ul",9),l.vc(8,y,9,3,"li",10),l.Ob(),l.vc(9,M,1,6,"ngb-pagination",11),l.Ob(),l.Ob()),2&e){const e=l.cc();l.xb(4),l.fc("routerLink",l.jc(4,C,e.usCurrentPage)),l.xb(2),l.fc("ngIf",e.ShowMessage),l.xb(2),l.fc("ngForOf",e.Users),l.xb(1),l.fc("ngIf",e.usCollectionSize>e.usPageSize)}}function I(e,t){1&e&&(l.Pb(0,"div",20),l.Pb(1,"span",21),l.xc(2,"Loading..."),l.Ob(),l.Ob())}let L=(()=>{class e{constructor(e,t,n,i){this.ActiveRoute=e,this.DataServ=t,this.AdminServ=n,this.router=i}ngOnDestroy(){this.RecivedErrorSub.unsubscribe(),this.PageChanged.unsubscribe(),this.DataLoading.unsubscribe(),this.FetchOnInint.unsubscribe()}ngOnInit(){this.usPageSize=b.a.AdminUserListPageSize,this.RecivedErrorSub=this.DataServ.RecivedError.subscribe(e=>{switch(this.ShowMessage=!0,this.ResponseFromBackend=e,setTimeout(()=>this.ShowMessage=!1,5e3),e.Error.Code){case 200:this.MessageType="success";break;default:this.MessageType="danger"}}),this.PageChanged=this.ActiveRoute.params.subscribe(e=>{this.usCurrentPage=+e.pn,this.FetchOnInint=this.DataServ.FetchUsersList(this.usCurrentPage,b.a.AdminUserListPageSize).subscribe(e=>{this.Users=this.AdminServ.GetUsers(),this.usCollectionSize=this.AdminServ.Total},e=>{this.Users=[]})}),this.DataLoading=this.DataServ.LoadingData.subscribe(e=>{this.IsLoading=e})}OnPageChanged(e){this.usCurrentPage=e,this.router.navigate(["../",e.toString()],{relativeTo:this.ActiveRoute})}OnDeleteUser(e,t){this.AdminServ.DeleteUser(t),this.DataServ.DeleteUser(e),this.Users=this.AdminServ.GetUsers()}}return e.\u0275fac=function(t){return new(t||e)(l.Kb(r.a),l.Kb(u.a),l.Kb(m.a),l.Kb(r.c))},e.\u0275cmp=l.Eb({type:e,selectors:[["app-user-list"]],decls:3,vars:2,consts:[["class","row",4,"ngIf"],[1,"text-center"],["class","spinner-border","role","status",4,"ngIf"],[1,"row"],[1,"col"],[1,"input-group","mb-1"],[1,"input-group-prepend","mt-1"],["queryParamsHandling","merge",1,"btn","btn-outline-primary",3,"routerLink"],[3,"type","close",4,"ngIf"],[1,"list-group","mb-1"],["style","cursor: pointer;","class","list-group-item list-group-item-action d-flex justify-content-between align-items-center",3,"ngClass","click",4,"ngFor","ngForOf"],[3,"collectionSize","pageSize","page","rotate","boundaryLinks","maxSize","pageChange",4,"ngIf"],[3,"type","close"],[1,"list-group-item","list-group-item-action","d-flex","justify-content-between","align-items-center",2,"cursor","pointer",3,"ngClass","click"],[1,"input-group"],[1,"p-2","m-0","flex-grow-1",2,"place-items","center"],[1,"input-group-append"],["type","button",1,"btn","btn-outline-success",3,"routerLink"],["type","button",1,"btn","btn-outline-danger",3,"click"],[3,"collectionSize","pageSize","page","rotate","boundaryLinks","maxSize","pageChange"],["role","status",1,"spinner-border"],[1,"sr-only"]],template:function(e,t){1&e&&(l.vc(0,k,10,6,"div",0),l.Pb(1,"div",1),l.vc(2,I,3,0,"div",2),l.Ob()),2&e&&(l.fc("ngIf",!t.IsLoading),l.xb(2),l.fc("ngIf",t.IsLoading))},directives:[i.k,r.d,i.j,o.a,i.i,o.i],styles:[""]}),e})();var U=n("Ja2n");function E(e,t){if(1&e){const e=l.Qb();l.Pb(0,"ngb-alert",11),l.ac("close",(function(){return l.oc(e),l.cc(2).ShowMessage=!1})),l.xc(1),l.Ob()}if(2&e){const e=l.cc(2);l.fc("type",e.MessageType),l.xb(1),l.zc(" ",e.ResponseFromBackend.Error.Message,"")}}function D(e,t){if(1&e){const e=l.Qb();l.Pb(0,"div",12),l.Pb(1,"a",13),l.Lb(2,"img",14),l.Ob(),l.Pb(3,"div",15),l.Pb(4,"button",16),l.xc(5,"Edit"),l.Ob(),l.Pb(6,"button",17),l.ac("click",(function(){l.oc(e);const n=t.$implicit,i=t.index;return l.cc(2).OnDeleteFile(n,i)})),l.xc(7,"Delete"),l.Ob(),l.Ob(),l.Ob()}if(2&e){const e=t.$implicit,n=t.index;l.xb(1),l.fc("href",e.FileID,l.pc),l.xb(1),l.fc("src",e.FileID,l.pc)("alt",e.Filename),l.xb(2),l.fc("routerLink",n)}}function z(e,t){if(1&e){const e=l.Qb();l.Pb(0,"ngb-pagination",18),l.ac("pageChange",(function(t){return l.oc(e),l.cc(2).OnPageChanged(t)})),l.Ob()}if(2&e){const e=l.cc(2);l.fc("collectionSize",e.meCollectionSize)("pageSize",e.mePageSize)("page",e.meCurrentPage)("rotate",!0)("boundaryLinks",!0)("maxSize",11)}}function T(e,t){if(1&e&&(l.Pb(0,"div"),l.vc(1,E,2,2,"ngb-alert",3),l.Pb(2,"div",4),l.Pb(3,"div",5),l.Pb(4,"button",6),l.xc(5,"Add"),l.Ob(),l.Ob(),l.Ob(),l.Pb(6,"div",7),l.vc(7,D,8,4,"div",8),l.Ob(),l.Pb(8,"div",9),l.vc(9,z,1,6,"ngb-pagination",10),l.Ob(),l.Ob()),2&e){const e=l.cc();l.xb(1),l.fc("ngIf",e.ShowMessage),l.xb(6),l.fc("ngForOf",e.Files),l.xb(2),l.fc("ngIf",e.meCollectionSize>e.mePageSize)}}function A(e,t){1&e&&(l.Pb(0,"div",19),l.Pb(1,"span",20),l.xc(2,"Loading..."),l.Ob(),l.Ob())}let F=(()=>{class e{constructor(e,t,n,i){this.ActiveRoute=e,this.DataServ=t,this.MediaServ=n,this.router=i}ngOnDestroy(){this.RecivedErrorSub.unsubscribe(),this.PageChanged.unsubscribe(),this.DataLoading.unsubscribe(),this.FetchOnInint.unsubscribe()}ngOnInit(){this.mePageSize=b.a.MediaListPageSize,this.RecivedErrorSub=this.DataServ.RecivedError.subscribe(e=>{switch(this.ShowMessage=!0,this.ResponseFromBackend=e,setTimeout(()=>this.ShowMessage=!1,5e3),e.Error.Code){case 200:this.MessageType="success";break;default:this.MessageType="danger"}}),this.PageChanged=this.ActiveRoute.params.subscribe(e=>{this.meCurrentPage=+e.pn,this.FetchOnInint=this.DataServ.FetchFilesList(this.meCurrentPage,b.a.MediaListPageSize).subscribe(e=>{this.Files=this.MediaServ.GetFiles(),this.meCollectionSize=this.MediaServ.Total},e=>{this.Files=[]})}),this.DataLoading=this.DataServ.LoadingData.subscribe(e=>{this.IsLoading=e})}OnPageChanged(e){this.meCurrentPage=e,this.router.navigate(["../",e.toString()],{relativeTo:this.ActiveRoute})}OnDeleteFile(e,t){this.MediaServ.DeleteFile(t),this.DataServ.DeleteFile(e.ID,!1),this.Files=this.MediaServ.GetFiles()}}return e.\u0275fac=function(t){return new(t||e)(l.Kb(r.a),l.Kb(u.a),l.Kb(U.a),l.Kb(r.c))},e.\u0275cmp=l.Eb({type:e,selectors:[["app-media-list"]],decls:3,vars:2,consts:[[4,"ngIf"],[1,"text-center"],["class","spinner-border","role","status",4,"ngIf"],[3,"type","close",4,"ngIf"],[1,"input-group","ml-3"],[1,"input-group-prepend","mt-2"],[1,"btn","btn-outline-primary"],[1,"parent-media-box"],["class","media-box",4,"ngFor","ngForOf"],[1,"ml-3"],[3,"collectionSize","pageSize","page","rotate","boundaryLinks","maxSize","pageChange",4,"ngIf"],[3,"type","close"],[1,"media-box"],[3,"href"],[3,"src","alt"],[1,"media-box-overlay"],["type","button",1,"btn","btn-outline-success",3,"routerLink"],["type","button",1,"btn","btn-outline-danger",3,"click"],[3,"collectionSize","pageSize","page","rotate","boundaryLinks","maxSize","pageChange"],["role","status",1,"spinner-border"],[1,"sr-only"]],template:function(e,t){1&e&&(l.vc(0,T,10,3,"div",0),l.Pb(1,"div",1),l.vc(2,A,3,0,"div",2),l.Ob()),2&e&&(l.fc("ngIf",!t.IsLoading),l.xb(2),l.fc("ngIf",t.IsLoading))},directives:[i.k,i.j,o.a,r.d,o.i],styles:[".parent-media-box[_ngcontent-%COMP%]{display:grid;grid-gap:1rem;grid-template-columns:repeat(auto-fit,minmax(360px,1fr));grid-auto-rows:minmax(240px,auto);padding:.3rem 1rem}.media-box[_ngcontent-%COMP%]{font-size:2rem;padding:.3rem;border-radius:5px;display:inline-block;position:relative;border:1px solid rgba(0,0,0,.125);color:#495057;width:100%;height:100%;-o-object-fit:cover;object-fit:cover}.media-box[_ngcontent-%COMP%]   img[_ngcontent-%COMP%]{max-width:100%;max-height:100%;border-radius:5px;display:block}.media-box-overlay[_ngcontent-%COMP%]{position:absolute;bottom:.5rem;right:.5rem}"]}),e})(),R=(()=>{class e{constructor(){}ngOnInit(){}}return e.\u0275fac=function(t){return new(t||e)},e.\u0275cmp=l.Eb({type:e,selectors:[["app-media-edit"]],decls:2,vars:0,template:function(e,t){1&e&&(l.Pb(0,"p"),l.xc(1,"media-edit works!"),l.Ob())},styles:[""]}),e})(),_=(()=>{class e{constructor(){}ngOnInit(){}}return e.\u0275fac=function(t){return new(t||e)},e.\u0275cmp=l.Eb({type:e,selectors:[["app-sessions-list"]],decls:2,vars:0,template:function(e,t){1&e&&(l.Pb(0,"p"),l.xc(1,"sessions-list works!"),l.Ob())},styles:[""]}),e})(),j=(()=>{class e{constructor(){}ngOnInit(){}}return e.\u0275fac=function(t){return new(t||e)},e.\u0275cmp=l.Eb({type:e,selectors:[["app-sessions-edit"]],decls:2,vars:0,template:function(e,t){1&e&&(l.Pb(0,"p"),l.xc(1,"sessions-edit works!"),l.Ob())},styles:[""]}),e})();const K=[{path:"",redirectTo:"users",pathMatch:"full"},{path:"users",component:g,canActivate:[a.a],children:[{path:"",redirectTo:"1",pathMatch:"full"},{path:":pn",component:L,canActivate:[c.a],data:{expectedRole:"admin_role_CRUD"}},{path:":pn/new",component:S},{path:":pn/:id",component:S,resolve:[p]}]},{path:"sessions",component:g,canActivate:[a.a],children:[{path:"",redirectTo:"1",pathMatch:"full"},{path:":pn",component:_,canActivate:[c.a],data:{expectedRole:"admin_role_CRUD"}},{path:":pn/new",component:j},{path:":pn/:id",component:j,resolve:[p]}]},{path:"media",component:g,canActivate:[a.a],children:[{path:"",redirectTo:"1",pathMatch:"full"},{path:":pn",component:F,canActivate:[c.a],data:{expectedRole:"admin_role_CRUD"}},{path:":pn/new",component:R},{path:":pn/:id",component:R,resolve:[p]}]}];let G=(()=>{class e{}return e.\u0275mod=l.Ib({type:e}),e.\u0275inj=l.Hb({factory:function(t){return new(t||e)},imports:[[r.g.forChild(K)],r.g]}),e})(),Q=(()=>{class e{}return e.\u0275mod=l.Ib({type:e}),e.\u0275inj=l.Hb({factory:function(t){return new(t||e)},imports:[[i.b,s.d,o.b,o.j,G]]}),e})()}}]);