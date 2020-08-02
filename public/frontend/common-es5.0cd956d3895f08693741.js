!function(){function e(e,t){for(var n=0;n<t.length;n++){var i=t[n];i.enumerable=i.enumerable||!1,i.configurable=!0,"value"in i&&(i.writable=!0),Object.defineProperty(e,i.key,i)}}function t(t,n,i){return n&&e(t.prototype,n),i&&e(t,i),t}function n(e,t){if(!(e instanceof t))throw new TypeError("Cannot call a class as a function")}(window.webpackJsonp=window.webpackJsonp||[]).push([[1],{EiSp:function(e,t,i){"use strict";i.d(t,"a",(function(){return r}));var r=function e(t,i){n(this,e),this.Name=t,this.Amount=i}},GXvH:function(e,i,r){"use strict";r.d(i,"a",(function(){return f}));var s=r("tk/3"),a=r("qc5x"),o=r("lJxs"),c=r("vkgz"),u=r("AytR"),d=r("XNiG"),h=r("fXoL"),p=r("ceC1"),l=r("ozzT"),g=r("l3fW"),f=function(){var e=function(){function e(t,i,r,s){n(this,e),this.http=t,this.recipes=i,this.shoppinglist=r,this.users=s,this.LoadingData=new d.a,this.RecipesUpdateInsert=new d.a,this.RecivedError=new d.a,this.PaginationSet=new d.a,this.FileUploadProgress=new d.a,this.FileUploaded=new d.a}return t(e,[{key:"FetchRecipes",value:function(e,t){var n=this;this.LoadingData.next(!0);var i={headers:new s.e({Page:e.toString(),Limit:t.toString()})};return this.http.get(u.a.GetSetRecipesUrl+"?key="+u.a.ApiKey,i).pipe(Object(o.a)((function(e){return e.Recipes=e.Recipes.map((function(e){return Object.assign(Object.assign({},e),{Ingredients:e.Ingredients?e.Ingredients:[]})})),e})),Object(c.a)((function(e){n.recipes.SetRecipes(e.Recipes),n.PaginationSet.next(new a.b(e.Total,e.Limit,e.Offset)),n.LoadingData.next(!1)}),(function(e){n.RecivedError.next(e.error),n.LoadingData.next(!1)})))}},{key:"SearchRecipes",value:function(e,t,n){var i=this;this.LoadingData.next(!0);var r={headers:new s.e({Page:e.toString(),Limit:t.toString(),Search:n})};return this.http.get(u.a.SearchRecipesUrl+"?key="+u.a.ApiKey,r).pipe(Object(o.a)((function(e){return e.Recipes=e.Recipes.map((function(e){return Object.assign(Object.assign({},e),{Ingredients:e.Ingredients?e.Ingredients:[]})})),e})),Object(c.a)((function(e){i.recipes.SetRecipes(e.Recipes),i.PaginationSet.next(new a.b(e.Total,e.Limit,e.Offset)),i.LoadingData.next(!1)}),(function(e){i.RecivedError.next(e.error),i.LoadingData.next(!1)})))}},{key:"SaveRecipe",value:function(e){var t=this;this.LoadingData.next(!0),this.http.post(u.a.GetSetRecipesUrl+"?key="+u.a.ApiKey,e).subscribe((function(e){t.RecipesUpdateInsert.next(e),t.RecivedError.next(new a.a(200,"\u0414\u0430\u043d\u043d\u044b\u0435 \u0441\u043e\u0445\u0440\u0430\u043d\u0435\u043d\u044b")),t.LoadingData.next(!1)}),(function(e){t.RecivedError.next(e.error),t.LoadingData.next(!1)}))}},{key:"FileUpload",value:function(e){var t=this,n=new FormData;n.append("image",e,e.name),this.http.post(u.a.UploadFileUrl+"?key="+u.a.ApiKey,n,{reportProgress:!0,observe:"events"}).subscribe((function(e){e.type===s.d.UploadProgress?t.FileUploadProgress.next(String(e.loaded/e.total*100)):e.type===s.d.Response&&e.ok&&t.FileUploaded.next(e.body)}),(function(e){t.RecivedError.next(e.error),t.LoadingData.next(!1)}))}},{key:"DeleteRecipe",value:function(e){var t=this;this.LoadingData.next(!0);var n={headers:new s.e({RecipeID:e.ID.toString()})};this.http.delete(u.a.GetSetRecipesUrl+"?key="+u.a.ApiKey,n).subscribe((function(e){t.RecivedError.next(e),t.LoadingData.next(!1)}),(function(e){t.RecivedError.next(e.error),t.LoadingData.next(!1)}))}},{key:"FetchShoppingList",value:function(e,t){var n=this;this.LoadingData.next(!0);var i={headers:new s.e({Page:e.toString(),Limit:t.toString()})};return this.http.get(u.a.GetSetShoppingListUrl+"?key="+u.a.ApiKey,i).pipe(Object(c.a)((function(e){n.shoppinglist.SetIngredients(e.Items),n.shoppinglist.SetPagination(e.Total,e.Limit,e.Offset),n.LoadingData.next(!1)}),(function(e){n.RecivedError.next(e.error),n.LoadingData.next(!1)})))}},{key:"SaveShoppingList",value:function(e){var t=this;this.LoadingData.next(!0),this.http.post(u.a.GetSetShoppingListUrl+"?key="+u.a.ApiKey,e).subscribe((function(e){t.RecipesUpdateInsert.next(e),t.RecivedError.next(new a.a(200,"\u0414\u0430\u043d\u043d\u044b\u0435 \u0441\u043e\u0445\u0440\u0430\u043d\u0435\u043d\u044b")),t.LoadingData.next(!1)}),(function(e){t.RecivedError.next(e.error),t.LoadingData.next(!1)}))}},{key:"DeleteShoppingList",value:function(e){var t=this;this.LoadingData.next(!0);var n={headers:new s.e({IngName:encodeURI(e.Name)})};this.http.delete(u.a.GetSetShoppingListUrl+"?key="+u.a.ApiKey,n).subscribe((function(e){t.RecivedError.next(e),t.LoadingData.next(!1)}),(function(e){t.RecivedError.next(e.error),t.LoadingData.next(!1)}))}},{key:"DeleteAllShoppingList",value:function(){var e=this;this.http.delete(u.a.GetSetShoppingListUrl+"?key="+u.a.ApiKey).subscribe((function(t){e.RecivedError.next(t),e.LoadingData.next(!1)}),(function(t){e.RecivedError.next(t.error),e.LoadingData.next(!1)}))}},{key:"FetchUsersList",value:function(e,t){var n=this;this.LoadingData.next(!0);var i={headers:new s.e({Page:e.toString(),Limit:t.toString()})};return this.http.get(u.a.GetSetUsersUrl+"?key="+u.a.ApiKey,i).pipe(Object(c.a)((function(e){n.users.SetUsers(e.Users),n.users.SetPagination(e.Total,e.Limit,e.Offset),n.LoadingData.next(!1)}),(function(e){n.RecivedError.next(e.error),n.LoadingData.next(!1)})))}}]),e}();return e.\u0275fac=function(t){return new(t||e)(h.Xb(s.b),h.Xb(p.a),h.Xb(l.a),h.Xb(g.a))},e.\u0275prov=h.Gb({token:e,factory:e.\u0275fac,providedIn:"root"}),e}()},ceC1:function(e,i,r){"use strict";r.d(i,"a",(function(){return h}));var s=r("qc5x"),a=r("EiSp"),o=r("XNiG"),c=r("AytR"),u=r("fXoL"),d=r("ozzT"),h=function(){var e=function(){function e(t){n(this,e),this.ShopList=t,this.IngredientSelected=new o.a,this.RecipeChanged=new o.a,this.RecipesUpdated=new o.a,this.RecipesInserted=new o.a,this.RecipesDeleted=new o.a,this.recipes=[]}return t(e,[{key:"GetRecipes",value:function(){return this.recipes.slice(0,c.a.RecipePageSize)}},{key:"GetRecipesLen",value:function(){return this.recipes.length}},{key:"GetRecipeById",value:function(e){return e<this.recipes.length&&e>0?this.recipes[e]:this.recipes[0]}},{key:"GetRecipeId",value:function(e){return this.recipes.indexOf(e)}},{key:"SendToShoppingList",value:function(e){var t=this;e.forEach((function(e){t.ShopList.AddNewItem(e)}))}},{key:"GetShoppingList",value:function(){return this.ShopList.GetIngredients()}},{key:"AddNewRecipe",value:function(e){var t=new s.c(e.Name,e.Description,e.ImagePath,e.Ingredients,e.ImageDbID,e.ID);this.recipes.length<c.a.RecipePageSize&&this.recipes.push(t),this.RecipeChanged.next(t),this.RecipesInserted.next()}},{key:"UpdateExistingRecipe",value:function(e,t){this.recipes[t]=e,this.RecipeChanged.next(e)}},{key:"SetRecipes",value:function(e){this.recipes=e,this.RecipesUpdated.next()}},{key:"DeleteRecipe",value:function(e){this.recipes.splice(e,1),this.RecipesDeleted.next()}},{key:"AddNewIngredient",value:function(e){var t=this.RecipeToEdit.Ingredients.find((function(t){return t.Name===e.Name}));t?t.Amount=t.Amount+e.Amount:this.RecipeToEdit.Ingredients.push(new a.a(e.Name,e.Amount))}},{key:"UpdateSelectedIngredient",value:function(e){var t=this,n=this.RecipeToEdit.Ingredients.find((function(t){return t.Name===e.Name}));if(n&&n!==this.CurrentSelectedItem)n.Name=e.Name,n.Amount=n.Amount+e.Amount,this.DeleteSelectedIngredient();else{var i=this.RecipeToEdit.Ingredients.find((function(e){return e===t.CurrentSelectedItem}));i&&(i.Name=e.Name,i.Amount=e.Amount)}this.CurrentSelectedItem=null,this.IngredientSelected.next(this.CurrentSelectedItem)}},{key:"DeleteSelectedIngredient",value:function(){var e=this.RecipeToEdit.Ingredients.indexOf(this.CurrentSelectedItem);-1!==e&&this.RecipeToEdit.Ingredients.splice(e,1),this.CurrentSelectedItem=null}},{key:"ClearAllIngredients",value:function(){this.RecipeToEdit.Ingredients=[],this.CurrentSelectedItem=null}},{key:"GetIngredientsLength",value:function(){return this.RecipeToEdit.Ingredients.length}},{key:"IngredientSelect",value:function(e){this.CurrentSelectedItem=e,this.IngredientSelected.next(e)}}]),e}();return e.\u0275fac=function(t){return new(t||e)(u.Xb(d.a))},e.\u0275prov=u.Gb({token:e,factory:e.\u0275fac,providedIn:"root"}),e}()},dxYa:function(e,i,r){"use strict";r.d(i,"a",(function(){return c}));var s=r("fXoL"),a=r("qXBG"),o=r("tyNb"),c=function(){var e=function(){function e(t,i){n(this,e),this.auth=t,this.router=i}return t(e,[{key:"canActivate",value:function(e,t){return e.data.expectedRole===this.auth.GetUserRole()||this.router.createUrlTree(["/"])}}]),e}();return e.\u0275fac=function(t){return new(t||e)(s.Xb(a.a),s.Xb(o.c))},e.\u0275prov=s.Gb({token:e,factory:e.\u0275fac,providedIn:"root"}),e}()},l3fW:function(e,i,r){"use strict";r.d(i,"a",(function(){return o}));var s=r("XNiG"),a=r("fXoL"),o=function(){var e=function(){function e(){n(this,e),this.UserSelected=new s.a,this.UsersUpdated=new s.a,this.Users=[]}return t(e,[{key:"GetUsers",value:function(){return this.Users.slice()}},{key:"SetUsers",value:function(e){this.Users=e,this.UsersUpdated.next()}},{key:"SetPagination",value:function(e,t,n){this.Total=e}},{key:"SelectItemUsersList",value:function(e){this.CurrentSelectedItem=e,this.UserSelected.next(e)}},{key:"IsCurrentSelected",value:function(e){return this.CurrentSelectedItem===e}}]),e}();return e.\u0275fac=function(t){return new(t||e)},e.\u0275prov=a.Gb({token:e,factory:e.\u0275fac,providedIn:"root"}),e}()},ozzT:function(e,i,r){"use strict";r.d(i,"a",(function(){return u}));var s=r("EiSp"),a=r("XNiG"),o=r("AytR"),c=r("fXoL"),u=function(){var e=function(){function e(){n(this,e),this.IngredientSelected=new a.a,this.IngredientChanged=new a.a,this.IngredientsUpdated=new a.a,this.IngredientAdded=new a.a,this.IngredientUpdated=new a.a,this.IngredientDeleted=new a.a,this.IngredientClear=new a.a,this.ingredients=[]}return t(e,[{key:"GetIngredients",value:function(){return this.ingredients.slice()}},{key:"GetIngredientsLength",value:function(){return this.ingredients.length}},{key:"SetIngredients",value:function(e){this.ingredients=e,this.IngredientsUpdated.next()}},{key:"SetPagination",value:function(e,t,n){this.Total=e}},{key:"AddNewItem",value:function(e){var t=this.ingredients.find((function(t){return t.Name===e.Name}));t?t.Amount=t.Amount+e.Amount:(t=e,this.ingredients.length<o.a.ShoppingListPageSize&&this.ingredients.push(new s.a(e.Name,e.Amount))),this.IngredientAdded.next(t),this.IngredientChanged.next(this.ingredients.slice())}},{key:"UpdateSelectedItem",value:function(e){var t=this.ingredients.find((function(t){return t.Name===e.Name}));if(t&&t!==this.CurrentSelectedItem)t.Name=e.Name,t.Amount=t.Amount+e.Amount,this.DeleteSelectedItem();else{var n=this.ingredients.indexOf(this.CurrentSelectedItem);-1!==n&&(this.ingredients[n]=e,this.IngredientChanged.next(this.ingredients.slice())),t=this.CurrentSelectedItem}this.IngredientUpdated.next(t),this.CurrentSelectedItem=null}},{key:"DeleteSelectedItem",value:function(){var e=this.ingredients.indexOf(this.CurrentSelectedItem);-1!==e&&this.ingredients.splice(e,1),this.IngredientDeleted.next(this.CurrentSelectedItem),this.IngredientChanged.next(this.ingredients.slice()),this.CurrentSelectedItem=null}},{key:"ClearAll",value:function(){this.ingredients=[],this.CurrentSelectedItem=null,this.IngredientChanged.next(this.ingredients.slice()),this.SetPagination(0,0,0),this.IngredientClear.next()}},{key:"SelectItemShopList",value:function(e){this.CurrentSelectedItem=e,this.IngredientSelected.next(e)}},{key:"IsCurrentSelected",value:function(e){return this.CurrentSelectedItem===e}}]),e}();return e.\u0275fac=function(t){return new(t||e)},e.\u0275prov=c.Gb({token:e,factory:e.\u0275fac,providedIn:"root"}),e}()},qc5x:function(e,t,i){"use strict";i.d(t,"c",(function(){return r})),i.d(t,"a",(function(){return s})),i.d(t,"b",(function(){return o}));var r=function e(t,i,r,s,a,o){n(this,e),this.Name=t,this.Description=i,this.ImagePath=r,this.Ingredients=s,this.ImageDbID=a,this.ID=o},s=function e(t,i){n(this,e),this.Error=new a(t,i)},a=function e(t,i){n(this,e),this.Code=t,this.Message=i},o=function e(t,i,r){n(this,e),this.Total=t,this.Limit=i,this.Offset=r}}}])}();