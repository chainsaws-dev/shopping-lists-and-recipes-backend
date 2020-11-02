!function(){function e(e,t){for(var n=0;n<t.length;n++){var i=t[n];i.enumerable=i.enumerable||!1,i.configurable=!0,"value"in i&&(i.writable=!0),Object.defineProperty(e,i.key,i)}}function t(t,n,i){return n&&e(t.prototype,n),i&&e(t,i),t}function n(e,t){if(!(e instanceof t))throw new TypeError("Cannot call a class as a function")}(window.webpackJsonp=window.webpackJsonp||[]).push([[0],{"/1gQ":function(e,t,i){"use strict";i.d(t,"b",(function(){return r})),i.d(t,"a",(function(){return s})),i.d(t,"c",(function(){return o}));var r=function e(t,i){n(this,e),this.Name=t,this.Amount=i},s=function e(t,i){n(this,e),this.Error=new a(t,i)},a=function e(t,i){n(this,e),this.Code=t,this.Message=i},o=function e(t,i,r){n(this,e),this.Total=t,this.Limit=i,this.Offset=r}},GXvH:function(e,i,r){"use strict";r.d(i,"a",(function(){return S}));var s=r("tk/3"),a=r("/1gQ"),o=r("lJxs"),c=r("vkgz"),u=r("AytR"),d=r("XNiG"),h=r("fXoL"),l=r("ceC1"),f=r("ozzT"),g=r("c7dm"),v=r("Ja2n"),p=r("uCSH"),S=function(){var e=function(){function e(t,i,r,s,a,o){n(this,e),this.http=t,this.recipes=i,this.shoppinglist=r,this.users=s,this.media=a,this.sessions=o,this.LoadingData=new d.a,this.RecipesUpdateInsert=new d.a,this.RecivedError=new d.a,this.PaginationSet=new d.a,this.FileUploadProgress=new d.a,this.FileUploaded=new d.a,this.UserUpdateInsert=new d.a,this.CurrentUserFetch=new d.a,this.TwoFactorSub=new d.a}return t(e,[{key:"FetchRecipes",value:function(e,t){var n=this;this.LoadingData.next(!0);var i={headers:new s.e({Page:e.toString(),Limit:t.toString()})};return this.http.get(u.a.GetSetRecipesUrl,i).pipe(Object(o.a)((function(e){return e.Recipes=e.Recipes.map((function(e){return Object.assign(Object.assign({},e),{Ingredients:e.Ingredients?e.Ingredients:[]})})),e})),Object(c.a)((function(e){n.recipes.SetRecipes(e.Recipes),n.PaginationSet.next(new a.c(e.Total,e.Limit,e.Offset)),n.LoadingData.next(!1)}),(function(e){n.RecivedError.next(e.error),n.LoadingData.next(!1)})))}},{key:"SearchRecipes",value:function(e,t,n){var i=this;this.LoadingData.next(!0);var r={headers:new s.e({Page:e.toString(),Limit:t.toString(),Search:n})};return this.http.get(u.a.SearchRecipesUrl,r).pipe(Object(o.a)((function(e){return e.Recipes=e.Recipes.map((function(e){return Object.assign(Object.assign({},e),{Ingredients:e.Ingredients?e.Ingredients:[]})})),e})),Object(c.a)((function(e){i.recipes.SetRecipes(e.Recipes),i.PaginationSet.next(new a.c(e.Total,e.Limit,e.Offset)),i.LoadingData.next(!1)}),(function(e){i.RecivedError.next(e.error),i.LoadingData.next(!1)})))}},{key:"SaveRecipe",value:function(e){var t=this;this.LoadingData.next(!0),this.http.post(u.a.GetSetRecipesUrl,e).subscribe((function(e){t.RecipesUpdateInsert.next(e),t.RecivedError.next(new a.a(200,"\u0414\u0430\u043d\u043d\u044b\u0435 \u0441\u043e\u0445\u0440\u0430\u043d\u0435\u043d\u044b")),t.LoadingData.next(!1)}),(function(e){t.RecivedError.next(e.error),t.LoadingData.next(!1)}))}},{key:"FetchFilesList",value:function(e,t){var n=this;this.LoadingData.next(!0);var i={headers:new s.e({Page:e.toString(),Limit:t.toString()})};return this.http.get(u.a.GetSetFileUrl,i).pipe(Object(c.a)((function(e){n.media.SetFiles(e.Files),n.media.SetPagination(e.Total,e.Limit,e.Offset),n.LoadingData.next(!1)}),(function(e){n.RecivedError.next(e.error),n.LoadingData.next(!1)})))}},{key:"FileUpload",value:function(e){var t=this,n=new FormData;n.append("file",e,e.name),this.http.post(u.a.GetSetFileUrl,n,{headers:new s.e({}),reportProgress:!0,observe:"events"}).subscribe((function(e){e.type===s.d.UploadProgress?t.FileUploadProgress.next(String(e.loaded/e.total*100)):e.type===s.d.Response&&e.ok&&t.FileUploaded.next(e.body)}),(function(e){t.RecivedError.next(e.error),t.LoadingData.next(!1)}))}},{key:"DeleteFile",value:function(e,t){var n=this;this.LoadingData.next(!0);var i={headers:new s.e({FileID:e.toString()})};this.http.delete(u.a.GetSetFileUrl,i).subscribe((function(e){t||n.RecivedError.next(e),n.LoadingData.next(!1)}),(function(e){n.RecivedError.next(e.error),n.LoadingData.next(!1)}))}},{key:"DeleteRecipe",value:function(e){var t=this;this.LoadingData.next(!0);var n={headers:new s.e({RecipeID:e.ID.toString()})};this.http.delete(u.a.GetSetRecipesUrl,n).subscribe((function(n){t.RecivedError.next(n),t.LoadingData.next(!1),t.DeleteFile(e.ImageDbID,!0)}),(function(e){t.RecivedError.next(e.error),t.LoadingData.next(!1)}))}},{key:"FetchShoppingList",value:function(e,t){var n=this;this.LoadingData.next(!0);var i={headers:new s.e({Page:e.toString(),Limit:t.toString()})};return this.http.get(u.a.GetSetShoppingListUrl,i).pipe(Object(c.a)((function(e){n.shoppinglist.SetIngredients(e.Items),n.shoppinglist.SetPagination(e.Total,e.Limit,e.Offset),n.LoadingData.next(!1)}),(function(e){n.RecivedError.next(e.error),n.LoadingData.next(!1)})))}},{key:"SaveShoppingList",value:function(e){var t=this;this.LoadingData.next(!0),this.http.post(u.a.GetSetShoppingListUrl,e).subscribe((function(e){t.RecipesUpdateInsert.next(e),t.RecivedError.next(new a.a(200,"\u0414\u0430\u043d\u043d\u044b\u0435 \u0441\u043e\u0445\u0440\u0430\u043d\u0435\u043d\u044b")),t.LoadingData.next(!1)}),(function(e){t.RecivedError.next(e.error),t.LoadingData.next(!1)}))}},{key:"DeleteShoppingList",value:function(e){var t=this;this.LoadingData.next(!0);var n={headers:new s.e({IngName:encodeURI(e.Name)})};this.http.delete(u.a.GetSetShoppingListUrl,n).subscribe((function(e){t.RecivedError.next(e),t.LoadingData.next(!1)}),(function(e){t.RecivedError.next(e.error),t.LoadingData.next(!1)}))}},{key:"DeleteAllShoppingList",value:function(){var e=this;this.http.delete(u.a.GetSetShoppingListUrl).subscribe((function(t){e.RecivedError.next(t),e.LoadingData.next(!1)}),(function(t){e.RecivedError.next(t.error),e.LoadingData.next(!1)}))}},{key:"FetchSessionsList",value:function(e,t){var n=this;this.LoadingData.next(!0);var i={headers:new s.e({Page:e.toString(),Limit:t.toString()})};return this.http.get(u.a.GetSetSessionsUrl,i).pipe(Object(c.a)((function(e){n.sessions.SetSessions(e.Sessions),n.sessions.SetPagination(e.Total,e.Limit,e.Offset),n.LoadingData.next(!1)}),(function(e){n.RecivedError.next(e.error),n.LoadingData.next(!1)})))}},{key:"DeleteSessionByToken",value:function(e){var t=this;this.LoadingData.next(!0);var n={headers:new s.e({Token:e})};this.http.delete(u.a.GetSetSessionsUrl,n).subscribe((function(e){t.RecivedError.next(e),t.LoadingData.next(!1)}),(function(e){t.RecivedError.next(e.error),t.LoadingData.next(!1)}))}},{key:"DeleteSessionByEmail",value:function(e){var t=this;this.LoadingData.next(!0);var n={headers:new s.e({Email:e})};this.http.delete(u.a.GetSetSessionsUrl,n).subscribe((function(e){t.RecivedError.next(e),t.LoadingData.next(!1)}),(function(e){t.RecivedError.next(e.error),t.LoadingData.next(!1)}))}},{key:"FetchUsersList",value:function(e,t){var n=this;this.LoadingData.next(!0);var i={headers:new s.e({Page:e.toString(),Limit:t.toString()})};return this.http.get(u.a.GetSetUsersUrl,i).pipe(Object(c.a)((function(e){n.users.SetUsers(e.Users),n.users.SetPagination(e.Total,e.Limit,e.Offset),n.LoadingData.next(!1)}),(function(e){n.RecivedError.next(e.error),n.LoadingData.next(!1)})))}},{key:"FetchCurrentUser",value:function(){var e=this;return this.LoadingData.next(!0),this.http.get(u.a.GetSetCurrentUserUrl).subscribe((function(t){e.CurrentUserFetch.next(t),e.LoadingData.next(!1)}),(function(t){e.RecivedError.next(t.error),e.LoadingData.next(!1)}))}},{key:"SaveCurrentUser",value:function(e,t,n){var i=this;this.LoadingData.next(!0),0===e.GUID.length&&(e.GUID="00000000-0000-0000-0000-000000000000"),this.GetObsForSaveCurrentUser(e,t,n).subscribe((function(e){i.UserUpdateInsert.next(e),i.RecivedError.next(new a.a(200,"\u0414\u0430\u043d\u043d\u044b\u0435 \u0441\u043e\u0445\u0440\u0430\u043d\u0435\u043d\u044b")),i.LoadingData.next(!1)}),(function(e){i.RecivedError.next(e.error),i.LoadingData.next(!1)}))}},{key:"SaveUser",value:function(e,t,n){var i=this;this.LoadingData.next(!0),0===e.GUID.length&&(e.GUID="00000000-0000-0000-0000-000000000000"),this.GetObsForSaveUser(e,t,n).subscribe((function(e){i.UserUpdateInsert.next(e),i.RecivedError.next(new a.a(200,"\u0414\u0430\u043d\u043d\u044b\u0435 \u0441\u043e\u0445\u0440\u0430\u043d\u0435\u043d\u044b")),i.LoadingData.next(!1)}),(function(e){i.RecivedError.next(e.error),i.LoadingData.next(!1)}))}},{key:"LinkTwoFactor",value:function(e,t){var n=this,i={headers:new s.e({Passcode:e})};this.http.post(u.a.TOTPSettingsUrl,t,i).subscribe((function(e){t.SecondFactor=!0,n.TwoFactorSub.next(t),n.RecivedError.next(e),n.LoadingData.next(!1)}),(function(e){n.RecivedError.next(e.error),n.LoadingData.next(!1)}))}},{key:"UnlinkTwoFactor",value:function(e){var t=this;this.http.delete(u.a.TOTPSettingsUrl).subscribe((function(n){e.SecondFactor=!1,t.TwoFactorSub.next(e),t.RecivedError.next(n),t.LoadingData.next(!1)}),(function(e){t.RecivedError.next(e.error),t.LoadingData.next(!1)}))}},{key:"GetObsForSaveCurrentUser",value:function(e,t,n){if(t){var i={headers:new s.e({NewPassword:encodeURI(n)})};return this.http.post(u.a.GetSetCurrentUserUrl,e,i)}return this.http.post(u.a.GetSetCurrentUserUrl,e)}},{key:"GetObsForSaveUser",value:function(e,t,n){if(t){var i={headers:new s.e({NewPassword:encodeURI(n)})};return this.http.post(u.a.GetSetUsersUrl,e,i)}return this.http.post(u.a.GetSetUsersUrl,e)}},{key:"DeleteUser",value:function(e){var t=this;this.LoadingData.next(!0);var n={headers:new s.e({UserID:encodeURI(e.GUID)})};this.http.delete(u.a.GetSetUsersUrl,n).subscribe((function(e){t.RecivedError.next(e),t.LoadingData.next(!1)}),(function(e){t.RecivedError.next(e.error),t.LoadingData.next(!1)}))}},{key:"ConfirmEmail",value:function(e){var t=this;this.LoadingData.next(!0);var n={headers:new s.e({Token:e})};this.http.post(u.a.ConfirmEmailUrl,null,n).subscribe((function(e){t.RecivedError.next(e),t.LoadingData.next(!1)}),(function(e){t.RecivedError.next(e.error),t.LoadingData.next(!1)}))}},{key:"SendEmailConfirmEmail",value:function(e){var t=this;this.LoadingData.next(!0);var n={headers:new s.e({Email:e})};this.http.post(u.a.ResendEmailUrl,null,n).subscribe((function(e){t.RecivedError.next(e),t.LoadingData.next(!1)}),(function(e){t.RecivedError.next(e.error),t.LoadingData.next(!1)}))}},{key:"SendEmailResetPassword",value:function(e){var t=this;this.LoadingData.next(!0);var n={headers:new s.e({Email:e})};this.http.post(u.a.SendEmailResetPassUrl,null,n).subscribe((function(e){t.RecivedError.next(e),t.LoadingData.next(!1)}),(function(e){t.RecivedError.next(e.error),t.LoadingData.next(!1)}))}},{key:"SubmitNewPassword",value:function(e,t){var n=this;this.LoadingData.next(!0);var i={headers:new s.e({Token:e,NewPassword:encodeURI(t)})};this.http.post(u.a.ResetPasswordUrl,null,i).subscribe((function(e){n.RecivedError.next(e),n.LoadingData.next(!1)}),(function(e){n.RecivedError.next(e.error),n.LoadingData.next(!1)}))}}]),e}();return e.\u0275fac=function(t){return new(t||e)(h.Zb(s.b),h.Zb(l.a),h.Zb(f.a),h.Zb(g.a),h.Zb(v.a),h.Zb(p.a))},e.\u0275prov=h.Ib({token:e,factory:e.\u0275fac,providedIn:"root"}),e}()},Ja2n:function(e,i,r){"use strict";r.d(i,"a",(function(){return u}));var s=r("vYfO"),a=r("XNiG"),o=r("AytR"),c=r("fXoL"),u=function(){var e=function(){function e(){n(this,e),this.FileSelected=new a.a,this.FilesUpdated=new a.a,this.FilesInserted=new a.a,this.FilesDeleted=new a.a,this.FilesChanged=new a.a,this.Files=[]}return t(e,[{key:"GetFiles",value:function(){return this.Files.slice()}},{key:"SetFiles",value:function(e){this.Files=e,this.FilesUpdated.next()}},{key:"SetPagination",value:function(e,t,n){this.Total=e}},{key:"SelectItemFilesList",value:function(e){this.CurrentSelectedItem=e,this.FileSelected.next(e)}},{key:"IsCurrentSelected",value:function(e){return this.CurrentSelectedItem===e}},{key:"GetFileById",value:function(e){return e<this.Files.length&&e>0?this.Files[e]:this.Files[0]}},{key:"UpdateExistingFile",value:function(e,t){this.Files[t]=e,this.FilesChanged.next(e)}},{key:"AddNewFile",value:function(e){var t=new s.a(e.Filename,e.Filesize,e.Filetype,e.FileID,e.ID);this.Files.length<o.a.AdminUserListPageSize&&this.Files.push(t),this.FilesChanged.next(t),this.FilesInserted.next()}},{key:"DeleteFile",value:function(e){this.Files.splice(e,1),this.FilesDeleted.next()}}]),e}();return e.\u0275fac=function(t){return new(t||e)},e.\u0275prov=c.Ib({token:e,factory:e.\u0275fac,providedIn:"root"}),e}()},c7dm:function(e,i,r){"use strict";r.d(i,"a",(function(){return u}));var s=r("ggGY"),a=r("XNiG"),o=r("AytR"),c=r("fXoL"),u=function(){var e=function(){function e(){n(this,e),this.UserSelected=new a.a,this.UsersUpdated=new a.a,this.UsersInserted=new a.a,this.UsersDeleted=new a.a,this.UsersChanged=new a.a,this.Users=[]}return t(e,[{key:"GetUsers",value:function(){return this.Users.slice()}},{key:"SetUsers",value:function(e){this.Users=e,this.UsersUpdated.next()}},{key:"SetPagination",value:function(e,t,n){this.Total=e}},{key:"SelectItemUsersList",value:function(e){this.CurrentSelectedItem=e,this.UserSelected.next(e)}},{key:"IsCurrentSelected",value:function(e){return this.CurrentSelectedItem===e}},{key:"GetUserById",value:function(e){return e<this.Users.length&&e>0?this.Users[e]:this.Users[0]}},{key:"UpdateExistingUser",value:function(e,t){this.Users[t]=e,this.UsersChanged.next(e)}},{key:"AddNewUser",value:function(e){var t=new s.a(e.Role,e.Email,e.Phone,e.Name);this.Users.length<o.a.AdminUserListPageSize&&this.Users.push(t),this.UsersChanged.next(t),this.UsersInserted.next()}},{key:"DeleteUser",value:function(e){this.Users.splice(e,1),this.UsersDeleted.next()}}]),e}();return e.\u0275fac=function(t){return new(t||e)},e.\u0275prov=c.Ib({token:e,factory:e.\u0275fac,providedIn:"root"}),e}()},ceC1:function(e,i,r){"use strict";r.d(i,"a",(function(){return h}));var s=r("qc5x"),a=r("/1gQ"),o=r("XNiG"),c=r("AytR"),u=r("fXoL"),d=r("ozzT"),h=function(){var e=function(){function e(t){n(this,e),this.ShopList=t,this.IngredientSelected=new o.a,this.RecipeChanged=new o.a,this.RecipesUpdated=new o.a,this.RecipesInserted=new o.a,this.RecipesDeleted=new o.a,this.recipes=[]}return t(e,[{key:"GetRecipes",value:function(){return this.recipes.slice(0,c.a.RecipePageSize)}},{key:"GetRecipesLen",value:function(){return this.recipes.length}},{key:"GetRecipeById",value:function(e){return e<this.recipes.length&&e>0?this.recipes[e]:this.recipes[0]}},{key:"GetRecipeId",value:function(e){return this.recipes.indexOf(e)}},{key:"SendToShoppingList",value:function(e){var t=this;e.forEach((function(e){t.ShopList.AddNewItem(e,!0)}))}},{key:"GetShoppingList",value:function(){return this.ShopList.GetIngredients()}},{key:"AddNewRecipe",value:function(e){var t=new s.a(e.Name,e.Description,e.ImagePath,e.Ingredients,e.ImageDbID,e.ID);this.recipes.length<c.a.RecipePageSize&&this.recipes.push(t),this.RecipeChanged.next(t),this.RecipesInserted.next()}},{key:"UpdateExistingRecipe",value:function(e,t){this.recipes[t]=e,this.RecipeChanged.next(e)}},{key:"SetRecipes",value:function(e){this.recipes=e,this.RecipesUpdated.next()}},{key:"DeleteRecipe",value:function(e){this.recipes.splice(e,1),this.RecipesDeleted.next()}},{key:"AddNewIngredient",value:function(e){var t=this.RecipeToEdit.Ingredients.find((function(t){return t.Name===e.Name}));t?t.Amount=t.Amount+e.Amount:this.RecipeToEdit.Ingredients.push(new a.b(e.Name,e.Amount))}},{key:"UpdateSelectedIngredient",value:function(e){var t=this,n=this.RecipeToEdit.Ingredients.find((function(t){return t.Name===e.Name}));if(n&&n!==this.CurrentSelectedItem)n.Name=e.Name,n.Amount=n.Amount+e.Amount,this.DeleteSelectedIngredient();else{var i=this.RecipeToEdit.Ingredients.find((function(e){return e===t.CurrentSelectedItem}));i&&(i.Name=e.Name,i.Amount=e.Amount)}this.CurrentSelectedItem=null,this.IngredientSelected.next(this.CurrentSelectedItem)}},{key:"DeleteSelectedIngredient",value:function(){var e=this.RecipeToEdit.Ingredients.indexOf(this.CurrentSelectedItem);-1!==e&&this.RecipeToEdit.Ingredients.splice(e,1),this.CurrentSelectedItem=null}},{key:"ClearAllIngredients",value:function(){this.RecipeToEdit.Ingredients=[],this.CurrentSelectedItem=null}},{key:"GetIngredientsLength",value:function(){return this.RecipeToEdit.Ingredients.length}},{key:"IngredientSelect",value:function(e){this.CurrentSelectedItem=e,this.IngredientSelected.next(e)}}]),e}();return e.\u0275fac=function(t){return new(t||e)(u.Zb(d.a))},e.\u0275prov=u.Ib({token:e,factory:e.\u0275fac,providedIn:"root"}),e}()},ggGY:function(e,t,i){"use strict";i.d(t,"a",(function(){return r}));var r=function e(t,i,r,s){n(this,e),this.GUID="",this.Role=t,this.Email=i,this.Phone=r,this.Name=s,this.IsAdmin=!1,this.Confirmed=!1,this.SecondFactor=!1,this.Disabled=!1}},ozzT:function(e,i,r){"use strict";r.d(i,"a",(function(){return u}));var s=r("/1gQ"),a=r("XNiG"),o=r("AytR"),c=r("fXoL"),u=function(){var e=function(){function e(){n(this,e),this.IngredientSelected=new a.a,this.IngredientChanged=new a.a,this.IngredientsUpdated=new a.a,this.IngredientAdded=new a.a,this.IngredientUpdated=new a.a,this.IngredientDeleted=new a.a,this.IngredientClear=new a.a,this.ingredients=[]}return t(e,[{key:"GetIngredients",value:function(){return this.ingredients.slice()}},{key:"GetIngredientsLength",value:function(){return this.ingredients.length}},{key:"SetIngredients",value:function(e){this.ingredients=e,this.IngredientsUpdated.next()}},{key:"SetPagination",value:function(e,t,n){this.Total=e}},{key:"AddNewItem",value:function(e,t){var n=this.ingredients.find((function(t){return t.Name===e.Name}));n?n.Amount=n.Amount+e.Amount:(n=e,(this.ingredients.length<o.a.ShoppingListPageSize||t)&&this.ingredients.push(new s.b(e.Name,e.Amount))),this.IngredientAdded.next(n),this.IngredientChanged.next(this.ingredients.slice())}},{key:"UpdateSelectedItem",value:function(e){var t=this.ingredients.find((function(t){return t.Name===e.Name}));if(t&&t!==this.CurrentSelectedItem)t.Name=e.Name,t.Amount=t.Amount+e.Amount,this.DeleteSelectedItem();else{var n=this.ingredients.indexOf(this.CurrentSelectedItem);-1!==n&&(this.ingredients[n]=e,this.IngredientChanged.next(this.ingredients.slice())),t=this.CurrentSelectedItem}this.IngredientUpdated.next(t),this.CurrentSelectedItem=null}},{key:"DeleteSelectedItem",value:function(){var e=this.ingredients.indexOf(this.CurrentSelectedItem);-1!==e&&this.ingredients.splice(e,1),this.IngredientDeleted.next(this.CurrentSelectedItem),this.IngredientChanged.next(this.ingredients.slice()),this.CurrentSelectedItem=null}},{key:"ClearAll",value:function(){this.ingredients=[],this.CurrentSelectedItem=null,this.IngredientChanged.next(this.ingredients.slice()),this.SetPagination(0,0,0),this.IngredientClear.next()}},{key:"SelectItemShopList",value:function(e){this.CurrentSelectedItem=e,this.IngredientSelected.next(e)}},{key:"IsCurrentSelected",value:function(e){return this.CurrentSelectedItem===e}}]),e}();return e.\u0275fac=function(t){return new(t||e)},e.\u0275prov=c.Ib({token:e,factory:e.\u0275fac,providedIn:"root"}),e}()},qc5x:function(e,t,i){"use strict";i.d(t,"a",(function(){return r}));var r=function e(t,i,r,s,a,o){n(this,e),this.Name=t,this.Description=i,this.ImagePath=r,this.Ingredients=s,this.ImageDbID=a,this.ID=o}},uCSH:function(e,i,r){"use strict";r.d(i,"a",(function(){return o}));var s=r("XNiG"),a=r("fXoL"),o=function(){var e=function(){function e(){n(this,e),this.SessionsUpdated=new s.a,this.SessionsInserted=new s.a,this.SessionsDeleted=new s.a,this.SessionsChanged=new s.a,this.SessionsSelected=new s.a,this.Sessions=[]}return t(e,[{key:"GetSessions",value:function(){return this.Sessions.slice()}},{key:"SetSessions",value:function(e){this.Sessions=e,this.SessionsUpdated.next()}},{key:"SetPagination",value:function(e,t,n){this.Total=e}},{key:"SelectItemSessionsList",value:function(e){this.CurrentSelectedItem=e,this.SessionsSelected.next(e)}},{key:"IsCurrentSelected",value:function(e){return this.CurrentSelectedItem===e}},{key:"GetSessionById",value:function(e){return e<this.Sessions.length&&e>0?this.Sessions[e]:this.Sessions[0]}},{key:"DeleteSession",value:function(e){this.Sessions.splice(e,1),this.SessionsDeleted.next()}}]),e}();return e.\u0275fac=function(t){return new(t||e)},e.\u0275prov=a.Ib({token:e,factory:e.\u0275fac,providedIn:"root"}),e}()},vYfO:function(e,t,i){"use strict";i.d(t,"a",(function(){return r}));var r=function e(t,i,r,s,a){n(this,e),this.Filename=t,this.Filesize=i,this.Filetype=r,this.FileID=s,this.ID=a}}}])}();