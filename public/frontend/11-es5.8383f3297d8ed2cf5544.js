!function(){function e(e,t){if(!(e instanceof t))throw new TypeError("Cannot call a class as a function")}function t(e,t){for(var n=0;n<t.length;n++){var i=t[n];i.enumerable=i.enumerable||!1,i.configurable=!0,"value"in i&&(i.writable=!0),Object.defineProperty(e,i.key,i)}}function n(e,n,i){return n&&t(e.prototype,n),i&&t(e,i),e}(window.webpackJsonp=window.webpackJsonp||[]).push([[11],{CXQP:function(t,i,s){"use strict";s.r(i),s.d(i,"ShoppingListModule",(function(){return w}));var o=s("ofXK"),r=s("AytR"),c=s("fXoL"),a=s("ozzT"),u=s("tyNb"),l=s("GXvH"),b=s("/1gQ"),d=s("3Pt+"),g=["f"];function p(e,t){if(1&e){var n=c.Qb();c.Pb(0,"button",14),c.ac("click",(function(){return c.oc(n),c.cc().DeleteSelectedItem()})),c.xc(1,"Delete"),c.Ob()}}function h(e,t){if(1&e){var n=c.Qb();c.Pb(0,"button",15),c.ac("click",(function(){return c.oc(n),c.cc().ClearAllItems()})),c.xc(1,"Clear"),c.Ob()}}var f,v=((f=function(){function t(n,i){e(this,t),this.ShopListServ=n,this.DataServ=i,this.editmode=!1}return n(t,[{key:"ngOnInit",value:function(){var e=this;this.ingselected=this.ShopListServ.IngredientSelected.subscribe((function(t){e.selectedingredient=t,e.editmode=!0,e.slEditForm.setValue({name:e.selectedingredient.Name,amount:e.selectedingredient.Amount})})),this.IngAdd=this.ShopListServ.IngredientAdded.subscribe((function(t){e.DataServ.SaveShoppingList(t)})),this.IngUpd=this.ShopListServ.IngredientUpdated.subscribe((function(t){e.DataServ.SaveShoppingList(t)})),this.IngDel=this.ShopListServ.IngredientDeleted.subscribe((function(t){e.DataServ.DeleteShoppingList(t)})),this.IngCle=this.ShopListServ.IngredientClear.subscribe((function(){e.DataServ.DeleteAllShoppingList()}))}},{key:"ngOnDestroy",value:function(){this.ingselected.unsubscribe(),this.IngAdd.unsubscribe(),this.IngUpd.unsubscribe(),this.IngDel.unsubscribe(),this.IngCle.unsubscribe()}},{key:"AddNewItem",value:function(e){if(e.valid){var t=e.value;this.ShopListServ.AddNewItem(new b.b(t.name,parseInt(t.amount,10)),!1)}}},{key:"UpdateItem",value:function(e){if(e.valid){var t=e.value;this.ShopListServ.UpdateSelectedItem(new b.b(t.name,parseInt(t.amount,10))),this.editmode=!1,this.slEditForm.reset()}}},{key:"DeleteSelectedItem",value:function(){this.ShopListServ.DeleteSelectedItem()}},{key:"ClearAllItems",value:function(){this.ShopListServ.ClearAll()}}]),t}()).\u0275fac=function(e){return new(e||f)(c.Kb(a.a),c.Kb(l.a))},f.\u0275cmp=c.Eb({type:f,selectors:[["app-shopping-edit"]],viewQuery:function(e,t){var n;1&e&&c.Bc(g,!0),2&e&&c.mc(n=c.bc())&&(t.slEditForm=n.first)},decls:18,vars:4,consts:[[1,"row"],[1,"col"],[3,"ngSubmit"],["f","ngForm"],[1,"input-group","mt-3"],[1,"col-sm-9","form-group"],["type","text","id","name","placeholder","Name","name","name","ngModel","","required","",1,"form-control"],[1,"col","form-group"],["type","number","id","amount","placeholder","Amount","name","amount","ngModel","","required","","pattern","^[1-9]+[0-9]*$",1,"form-control"],[1,"input-group"],[1,"input-group-prepend"],["type","submit",1,"btn","btn-outline-primary",3,"disabled"],["class","btn btn-outline-danger","type","button",3,"click",4,"ngIf"],["class","btn btn-outline-secondary","type","button",3,"click",4,"ngIf"],["type","button",1,"btn","btn-outline-danger",3,"click"],["type","button",1,"btn","btn-outline-secondary",3,"click"]],template:function(e,t){if(1&e){var n=c.Qb();c.Pb(0,"div",0),c.Pb(1,"div",1),c.Pb(2,"form",2,3),c.ac("ngSubmit",(function(){c.oc(n);var e=c.nc(3);return t.editmode?t.UpdateItem(e):t.AddNewItem(e)})),c.Pb(4,"div",0),c.Pb(5,"div",4),c.Pb(6,"div",5),c.Lb(7,"input",6),c.Ob(),c.Pb(8,"div",7),c.Lb(9,"input",8),c.Ob(),c.Ob(),c.Ob(),c.Pb(10,"div",0),c.Pb(11,"div",1),c.Pb(12,"div",9),c.Pb(13,"div",10),c.Pb(14,"button",11),c.xc(15),c.Ob(),c.vc(16,p,2,0,"button",12),c.vc(17,h,2,0,"button",13),c.Ob(),c.Ob(),c.Ob(),c.Ob(),c.Ob(),c.Ob(),c.Ob()}if(2&e){var i=c.nc(3);c.xb(14),c.fc("disabled",i.invalid),c.xb(1),c.yc(t.editmode?"Update":"Add"),c.xb(1),c.fc("ngIf",t.ShopListServ.CurrentSelectedItem),c.xb(1),c.fc("ngIf",0!==t.ShopListServ.GetIngredientsLength())}},directives:[d.r,d.i,d.j,d.b,d.h,d.k,d.o,d.m,d.n,o.k],styles:["input.ng-invalid.ng-touched[_ngcontent-%COMP%]{border-color:#ff0a1d!important;box-shadow:none!important}input.ng-invalid.ng-touched[_ngcontent-%COMP%]:focus{box-shadow:0 0 0 .2rem rgba(255,10,29,.25)!important}input.ng-valid.ng-touched[_ngcontent-%COMP%]{border-color:#67c07b!important;box-shadow:none!important}input.ng-valid.ng-touched[_ngcontent-%COMP%]:focus{box-shadow:0 0 0 .2rem rgba(52,179,13,.25)!important}"]}),f),S=s("1kSV");function m(e,t){if(1&e){var n=c.Qb();c.Pb(0,"ngb-alert",9),c.ac("close",(function(){return c.oc(n),c.cc(2).ShowMessage=!1})),c.xc(1),c.Ob()}if(2&e){var i=c.cc(2);c.fc("type",i.MessageType),c.xb(1),c.yc(i.ResponseFromBackend.Error.Message)}}function I(e,t){if(1&e){var n=c.Qb();c.Pb(0,"a",10),c.ac("click",(function(){c.oc(n);var e=t.$implicit;return c.cc(2).ShopListServ.SelectItemShopList(e)})),c.xc(1),c.Pb(2,"span",11),c.xc(3),c.Ob(),c.Ob()}if(2&e){var i=t.$implicit,s=c.cc(2);c.fc("ngClass",s.ShopListServ.IsCurrentSelected(i)?"active":""),c.xb(1),c.zc("",i.Name," "),c.xb(2),c.yc(i.Amount)}}function L(e,t){if(1&e){var n=c.Qb();c.Pb(0,"ngb-pagination",12),c.ac("pageChange",(function(e){return c.oc(n),c.cc(2).OnPageChanged(e)})),c.Ob()}if(2&e){var i=c.cc(2);c.fc("collectionSize",i.slcollectionSize)("pageSize",i.slPageSize)("page",i.slCurrentPage)("rotate",!0)("boundaryLinks",!0)("maxSize",11)}}function P(e,t){if(1&e&&(c.Pb(0,"div",3),c.Pb(1,"div",4),c.Lb(2,"app-shopping-edit"),c.Lb(3,"hr"),c.vc(4,m,2,2,"ngb-alert",5),c.Pb(5,"ul",6),c.vc(6,I,4,3,"a",7),c.Ob(),c.vc(7,L,1,6,"ngb-pagination",8),c.Ob(),c.Ob()),2&e){var n=c.cc();c.xb(4),c.fc("ngIf",n.ShowMessage),c.xb(2),c.fc("ngForOf",n.ingredients),c.xb(1),c.fc("ngIf",n.slcollectionSize>n.slPageSize)}}function y(e,t){1&e&&(c.Pb(0,"div",13),c.Pb(1,"span",14),c.xc(2,"Loading..."),c.Ob(),c.Ob())}var C,O,k=[{path:"",redirectTo:"1",pathMatch:"full"},{path:":pn",component:(C=function(){function t(n,i,s,o){e(this,t),this.ShopListServ=n,this.activeroute=i,this.DataServ=s,this.router=o}return n(t,[{key:"ngOnDestroy",value:function(){this.IngChanged.unsubscribe(),this.PageChanged.unsubscribe(),this.FetchOnInint.unsubscribe(),this.DataLoading.unsubscribe(),this.RecivedErrorSub.unsubscribe(),this.WatchIngAdd.unsubscribe(),this.WatchIngDel.unsubscribe(),this.WatchIngCle.unsubscribe()}},{key:"ngOnInit",value:function(){var e=this;this.slPageSize=r.a.ShoppingListPageSize,this.RecivedErrorSub=this.DataServ.RecivedError.subscribe((function(t){switch(e.ShowMessage=!0,e.ResponseFromBackend=t,setTimeout((function(){return e.ShowMessage=!1}),5e3),t.Error.Code){case 200:e.MessageType="success";break;default:e.MessageType="danger"}})),this.IngChanged=this.ShopListServ.IngredientChanged.subscribe((function(t){e.ingredients=t})),this.PageChanged=this.activeroute.params.subscribe((function(t){e.slCurrentPage=+t.pn})),this.DataLoading=this.DataServ.LoadingData.subscribe((function(t){e.IsLoading=t})),this.FetchOnInint=this.DataServ.FetchShoppingList(this.slCurrentPage,r.a.ShoppingListPageSize).subscribe((function(t){e.ingredients=e.ShopListServ.GetIngredients(),e.slcollectionSize=e.ShopListServ.Total}),(function(t){e.ingredients=[]})),this.WatchIngAdd=this.ShopListServ.IngredientAdded.subscribe((function(t){e.slcollectionSize+=1,e.ingredients=e.ShopListServ.GetIngredients()})),this.WatchIngDel=this.ShopListServ.IngredientDeleted.subscribe((function(t){e.slcollectionSize-=1,e.ingredients=e.ShopListServ.GetIngredients(),0===e.ingredients.length&&(e.slCurrentPage=e.GetPreviousPage(e.slCurrentPage),e.ShopListServ.Total=e.slcollectionSize,0!==e.slcollectionSize&&e.OnPageChanged(e.slCurrentPage))})),this.WatchIngCle=this.ShopListServ.IngredientClear.subscribe((function(){e.slcollectionSize=0,e.ShopListServ.Total=e.slcollectionSize}))}},{key:"GetPreviousPage",value:function(e){return e>1?e-1:1}},{key:"OnPageChanged",value:function(e){var t=this;this.slCurrentPage=e,this.FetchOnInint=this.DataServ.FetchShoppingList(e,r.a.ShoppingListPageSize).subscribe((function(){t.ingredients=t.ShopListServ.GetIngredients(),t.router.navigate(["../",e.toString()],{relativeTo:t.activeroute})}))}}]),t}(),C.\u0275fac=function(e){return new(e||C)(c.Kb(a.a),c.Kb(u.a),c.Kb(l.a),c.Kb(u.c))},C.\u0275cmp=c.Eb({type:C,selectors:[["app-shopping-list"]],decls:3,vars:2,consts:[["class","row",4,"ngIf"],[1,"text-center"],["class","spinner-border","role","status",4,"ngIf"],[1,"row"],[1,"col"],[3,"type","close",4,"ngIf"],[1,"list-group","mb-1"],["style","cursor: pointer;","class","list-group-item list-group-item-action d-flex justify-content-between align-items-center",3,"ngClass","click",4,"ngFor","ngForOf"],[3,"collectionSize","pageSize","page","rotate","boundaryLinks","maxSize","pageChange",4,"ngIf"],[3,"type","close"],[1,"list-group-item","list-group-item-action","d-flex","justify-content-between","align-items-center",2,"cursor","pointer",3,"ngClass","click"],[1,"badge","badge-success","badge-pill"],[3,"collectionSize","pageSize","page","rotate","boundaryLinks","maxSize","pageChange"],["role","status",1,"spinner-border"],[1,"sr-only"]],template:function(e,t){1&e&&(c.vc(0,P,8,3,"div",0),c.Pb(1,"div",1),c.vc(2,y,3,0,"div",2),c.Ob()),2&e&&(c.fc("ngIf",!t.IsLoading),c.xb(2),c.fc("ngIf",t.IsLoading))},directives:[o.k,v,o.j,S.a,o.i,S.i],styles:[""]}),C),canActivate:[s("dxYa").a],data:{expectedRole:"admin_role_CRUD"}}],w=((O=function t(){e(this,t)}).\u0275mod=c.Ib({type:O}),O.\u0275inj=c.Hb({factory:function(e){return new(e||O)},imports:[[o.b,d.d,S.b,S.j,u.g.forChild(k)]]}),O)}}])}();