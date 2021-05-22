(function(e){function t(t){for(var n,s,i=t[0],u=t[1],c=t[2],m=0,f=[];m<i.length;m++)s=i[m],Object.prototype.hasOwnProperty.call(o,s)&&o[s]&&f.push(o[s][0]),o[s]=0;for(n in u)Object.prototype.hasOwnProperty.call(u,n)&&(e[n]=u[n]);l&&l(t);while(f.length)f.shift()();return a.push.apply(a,c||[]),r()}function r(){for(var e,t=0;t<a.length;t++){for(var r=a[t],n=!0,i=1;i<r.length;i++){var u=r[i];0!==o[u]&&(n=!1)}n&&(a.splice(t--,1),e=s(s.s=r[0]))}return e}var n={},o={app:0},a=[];function s(t){if(n[t])return n[t].exports;var r=n[t]={i:t,l:!1,exports:{}};return e[t].call(r.exports,r,r.exports,s),r.l=!0,r.exports}s.m=e,s.c=n,s.d=function(e,t,r){s.o(e,t)||Object.defineProperty(e,t,{enumerable:!0,get:r})},s.r=function(e){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(e,"__esModule",{value:!0})},s.t=function(e,t){if(1&t&&(e=s(e)),8&t)return e;if(4&t&&"object"===typeof e&&e&&e.__esModule)return e;var r=Object.create(null);if(s.r(r),Object.defineProperty(r,"default",{enumerable:!0,value:e}),2&t&&"string"!=typeof e)for(var n in e)s.d(r,n,function(t){return e[t]}.bind(null,n));return r},s.n=function(e){var t=e&&e.__esModule?function(){return e["default"]}:function(){return e};return s.d(t,"a",t),t},s.o=function(e,t){return Object.prototype.hasOwnProperty.call(e,t)},s.p="/";var i=window["webpackJsonp"]=window["webpackJsonp"]||[],u=i.push.bind(i);i.push=t,i=i.slice();for(var c=0;c<i.length;c++)t(i[c]);var l=u;a.push([0,"chunk-vendors"]),r()})({0:function(e,t,r){e.exports=r("56d7")},"56d7":function(e,t,r){"use strict";r.r(t);r("e260"),r("e6cf"),r("cca6"),r("a79d");var n=r("2b0e"),o=function(){var e=this,t=e.$createElement,r=e._self._c||t;return r("div",{attrs:{id:"app"}},[r("div",{staticClass:"d-flex",attrs:{id:"nav"}},[r("router-link",{staticClass:"mr-auto p-2",attrs:{to:"/"}},[e._v("Главная")]),r("router-link",{directives:[{name:"show",rawName:"v-show",value:!e.isLoggedIn,expression:"!isLoggedIn"}],staticClass:"p-2",attrs:{to:"/sign-in"}},[e._v("Авторизироваться")]),r("router-link",{directives:[{name:"show",rawName:"v-show",value:!e.isLoggedIn,expression:"!isLoggedIn"}],staticClass:"p-2",attrs:{to:"/sign-up"}},[e._v("Зарегестрироваться")]),r("button",{directives:[{name:"show",rawName:"v-show",value:e.isLoggedIn,expression:"isLoggedIn"}],staticClass:" btn p-2",on:{click:e.logOut}},[e._v("Выход")])],1),r("router-view")],1)},a=[],s={name:"Home",data:function(){return{}},created:function(){var e=this.$cookies.get("userdata");null!=e&&this.$store.commit("setUser",e)},updated:function(){this.isLoggedIn||this.$router.push("/sign-in")},methods:{logOut:function(){this.$cookies.remove("userdata"),this.$store.commit("setUser",void 0),this.$router.push("/"),this.userData=void 0}},computed:{isLoggedIn:function(){return void 0!==this.$store.getters.getUserData},userData:function(){return this.$store.getters.getUserData}}},i=s,u=(r("5c0b"),r("2877")),c=Object(u["a"])(i,o,a,!1,null,null,null),l=c.exports,m=r("8c4f"),f=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",{staticClass:"home"},[n("img",{attrs:{alt:"Vue logo",src:r("cf05")}}),e._v(" "+e._s(e.userData.first_name)+" "+e._s(e.userData.second_name)+" ")])},d=[],p={name:"Home",data:function(){return{}},created:function(){void 0==this.$store.getters.getUserData&&this.$router.push("/sign-in")},computed:{userData:function(){return this.$store.getters.getUserData}}},b=p,g=Object(u["a"])(b,f,d,!1,null,null,null),v=g.exports,h=function(){var e=this,t=e.$createElement,r=e._self._c||t;return r("div",{staticClass:"w-100 d-flex align-items-center flex-column"},[r("div",{staticClass:"p-3 w-50"},[e.show?r("b-form",{attrs:{autocomplete:"off"},on:{submit:e.onSubmit,reset:e.onReset}},[r("b-alert",{attrs:{variant:"danger"},model:{value:e.wrongUsername,callback:function(t){e.wrongUsername=t},expression:"wrongUsername"}},[e._v(" Проверте введенные данные ")]),r("b-form-group",{attrs:{id:"login-group",label:"Логин:","label-for":"login-input"}},[r("b-form-input",{attrs:{id:"login-input",autocomplete:"off",placeholder:"Введите логин",type:"text",required:""},model:{value:e.form.username,callback:function(t){e.$set(e.form,"username",t)},expression:"form.username"}})],1),r("b-form-group",{attrs:{id:"password-group",label:"Пароль:","label-for":"password-input"}},[r("b-form-input",{attrs:{id:"password-input",type:"password",autocomplete:"new-password",placeholder:"Введите пароль",required:""},model:{value:e.form.password,callback:function(t){e.$set(e.form,"password",t)},expression:"form.password"}})],1),r("b-form-group",{attrs:{id:"input-group-4"},scopedSlots:e._u([{key:"default",fn:function(t){var n=t.ariaDescribedby;return[r("b-form-checkbox-group",{attrs:{id:"checkboxes-4","aria-describedby":n},model:{value:e.form.checked,callback:function(t){e.$set(e.form,"checked",t)},expression:"form.checked"}},[r("b-form-checkbox",{model:{value:e.rememberMe,callback:function(t){e.rememberMe=t},expression:"rememberMe"}},[e._v("Запомнить")])],1)]}}],null,!1,3976312111)}),r("b-button",{staticClass:"m-1",attrs:{type:"submit",variant:"primary"}},[e._v("Авторизироваться")]),r("br"),r("b-button",{staticClass:"m-1",attrs:{type:"reset",variant:"danger"}},[e._v("Зарегестрироваться")])],1):e._e()],1)])},w=[],_=r("1da1"),x=(r("96cf"),r("d3b7"),{data:function(){return{form:{username:"",password:""},wrongUsername:!1,codes:[{text:"Выберите подразделение",value:null},"4157 Биробиджанское отделение","9070 ГО по Хабаровскому краю","8635 Приморское отделение","8636 Благовещенское отделение","8567 Ю.-Сахалинское отделение","8645 С.-Восточное отделение","8557 Чукотское отделение","8556 Камчатское отделение"],vacancies:[{text:"Выберите вакансию",value:null},"СМО (Старший менеджер по обслуживанию)","ВМО (Ведущий менеджер по обслуживанию)","МО (Менеджер по обслуживанию)","СКМ (Старший клиентский менеджер)","КМ (клиентский менеджер)","К (консультант)"],show:!0,rememberMe:!1}},methods:{onSubmit:function(e){var t=this;return Object(_["a"])(regeneratorRuntime.mark((function r(){var n,o;return regeneratorRuntime.wrap((function(r){while(1)switch(r.prev=r.next){case 0:return e.preventDefault(),r.next=3,fetch("/sing_in",{method:"POST",headers:{"Content-Type":"application/json;charset=utf-8"},body:JSON.stringify(t.form)});case 3:return n=r.sent,r.next=6,n.json();case 6:o=r.sent,400==o.code&&(t.wrongUsername=!0),200==o.code&&(t.$store.commit("setUser",o.data),t.rememberMe&&t.$cookies.set("userdata",o.data),t.$router.push("/"));case 9:case"end":return r.stop()}}),r)})))()},onReset:function(){this.$router.push("/sign-up")}}}),y=x,$=Object(u["a"])(y,h,w,!1,null,null,null),k=$.exports,O=function(){var e=this,t=e.$createElement,r=e._self._c||t;return r("div",{staticClass:"w-100 d-flex align-items-center flex-column"},[r("div",{staticClass:"p-3 w-50"},[e.show?r("b-form",{attrs:{autocomplete:"off"},on:{submit:e.onSubmit,reset:e.onReset}},[r("b-alert",{attrs:{variant:"danger"},model:{value:e.wrongUsername,callback:function(t){e.wrongUsername=t},expression:"wrongUsername"}},[e._v(" Имя пользователя занято ")]),r("b-form-group",{attrs:{id:"login-group",label:"Логин:","label-for":"login-input"}},[r("b-form-input",{attrs:{id:"login-input",autocomplete:"off",placeholder:"Введите логин",type:"text",required:""},model:{value:e.form.username,callback:function(t){e.$set(e.form,"username",t)},expression:"form.username"}})],1),r("b-form-group",{attrs:{id:"password-group",label:"Пароль:","label-for":"password-input"}},[r("b-form-input",{attrs:{id:"password-input",type:"password",autocomplete:"new-password",placeholder:"Введите пароль",required:""},model:{value:e.form.password,callback:function(t){e.$set(e.form,"password",t)},expression:"form.password"}})],1),r("b-form-group",{attrs:{id:"firstname-group",label:"Имя:","label-for":"firstname-input"}},[r("b-form-input",{attrs:{id:"firstname-input",placeholder:"Введите имя",required:""},model:{value:e.form.first_name,callback:function(t){e.$set(e.form,"first_name",t)},expression:"form.first_name"}})],1),r("b-form-group",{attrs:{id:"secondname-group",label:"Фамилия:","label-for":"secondname-input"}},[r("b-form-input",{attrs:{id:"secondname-input",placeholder:"Введите фамилию",required:""},model:{value:e.form.second_name,callback:function(t){e.$set(e.form,"second_name",t)},expression:"form.second_name"}})],1),r("b-form-group",{attrs:{id:"vacancy-group",label:"Вакансия:","label-for":"vacancy-select"}},[r("b-form-select",{staticClass:"form-select",attrs:{id:"vacancy-select",options:e.vacancies,required:""},model:{value:e.form.vacancy,callback:function(t){e.$set(e.form,"vacancy",t)},expression:"form.vacancy"}})],1),r("b-form-group",{attrs:{id:"code-group",label:"Подразделение:","label-for":"code-select"}},[r("b-form-select",{staticClass:"form-select",attrs:{id:"code-select",options:e.codes,required:""},model:{value:e.form.code,callback:function(t){e.$set(e.form,"code",t)},expression:"form.code"}})],1),r("b-button",{staticClass:"m-1",attrs:{type:"submit",variant:"primary"}},[e._v("Зарегестрироваться")]),r("br"),r("b-button",{staticClass:"m-1",attrs:{type:"reset",variant:"danger"}},[e._v("Авторизироваться")])],1):e._e()],1)])},U=[],j=(r("ac1f"),r("1276"),{data:function(){return{form:{username:"",password:"",first_name:"",second_name:"",code:null,vacancy:null},wrongUsername:!1,codes:[{text:"Выберите подразделение",value:null},"4157 Биробиджанское отделение","9070 ГО по Хабаровскому краю","8635 Приморское отделение","8636 Благовещенское отделение","8567 Ю.-Сахалинское отделение","8645 С.-Восточное отделение","8557 Чукотское отделение","8556 Камчатское отделение"],vacancies:[{text:"Выберите вакансию",value:null},"СМО (Старший менеджер по обслуживанию)","ВМО (Ведущий менеджер по обслуживанию)","МО (Менеджер по обслуживанию)","СКМ (Старший клиентский менеджер)","КМ (клиентский менеджер)","К (консультант)"],show:!0}},methods:{onSubmit:function(e){var t=this;return Object(_["a"])(regeneratorRuntime.mark((function r(){var n,o,a;return regeneratorRuntime.wrap((function(r){while(1)switch(r.prev=r.next){case 0:return e.preventDefault(),n={},n.username=t.form.username,n.password=t.form.password,n.first_name=t.form.first_name,n.second_name=t.form.second_name,n.code=t.form.code.split(" ")[0],n.vacancy=t.form.vacancy.split(" ")[0],r.next=10,fetch("/sing_up",{method:"POST",headers:{"Content-Type":"application/json;charset=utf-8"},body:JSON.stringify(n)});case 10:return o=r.sent,r.next=13,o.json();case 13:a=r.sent,400==a.code&&(t.wrongUsername=!0),200==a.code&&(t.$store.commit("setUser",a.data),t.rememberMe&&t.$cookies.set("userdata",a.data),t.$router.push("/"));case 16:case"end":return r.stop()}}),r)})))()},onReset:function(){this.$router.push("/sign-in")}}}),C=j,D=Object(u["a"])(C,O,U,!1,null,null,null),S=D.exports;n["default"].use(m["a"]);var M=[{path:"/",name:"Home",component:v},{path:"/sign-in",name:"sign-in",component:k},{path:"/sign-up",name:"sign-up",component:S}],P=new m["a"]({mode:"history",base:"/",routes:M}),q=P,I=r("2f62");n["default"].use(I["a"]);var L=new I["a"].Store({state:{userData:void 0},mutations:{setUser:function(e,t){e.userData=t}},actions:{},modules:{},getters:{getUserData:function(e){return e.userData}}}),R=r("5f5b"),T=r("b1e0"),N=r("2b27"),E=r.n(N);r("f9e3"),r("2dd8");n["default"].use(R["a"]),n["default"].use(T["a"]),n["default"].use(E.a),n["default"].$cookies.config("7d"),n["default"].config.productionTip=!1,new n["default"]({router:q,store:L,render:function(e){return e(l)}}).$mount("#app")},"5c0b":function(e,t,r){"use strict";r("9c0c")},"9c0c":function(e,t,r){},cf05:function(e,t,r){e.exports=r.p+"img/logo.png"}});