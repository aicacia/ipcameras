import{n as Ne,r as ze,c as be,m as st,d as Ge}from"../chunks/scheduler.CgEFyQ-S.js";import{S as Be,i as Se,e as v,t as ee,s as O,c as ne,a as g,b as y,g as te,h as m,f as q,d as re,j as r,k as we,l as n,u as le,m as ae,v as G,w as Ke,n as se,o as M,x as Oe,p as W,y as qe,q as oe,r as lt,C as Fe,D as nt}from"../chunks/index.DfAfitrp.js";import{L as Ue}from"../chunks/i18n-svelte.DO6JJ3x7.js";import{I as _e,a as Qe,S as Xe,b as Ye,O as Ze,c as ve,v as ge,h as xe}from"../chunks/InputResults.C9hpT5ea.js";import{d as et}from"../chunks/debounce.DpwXS7NU.js";import{b as rt,d as tt,e as at,C as ot,f as ut,g as it,A as ct,h as Me}from"../chunks/index.ibyW_jwT.js";import{p as ft,s as pt,h as dt,b as ht}from"../chunks/httpAccess.oZ7xIxns.js";import{g as mt}from"../chunks/entry.C7xj7j3Z.js";import{b as _t}from"../chunks/paths.CDMbSYKj.js";function Je(e){let s,t,l,u;return l=new Xe({}),{c(){s=v("div"),t=v("div"),ne(l.$$.fragment),this.h()},l(a){s=g(a,"DIV",{class:!0});var p=y(s);t=g(p,"DIV",{class:!0});var o=y(t);re(l.$$.fragment,o),o.forEach(m),p.forEach(m),this.h()},h(){r(t,"class","inline-block h-6 w-6"),r(s,"class","mr-2 flex flex-row justify-center")},m(a,p){we(a,s,p),n(s,t),ae(l,t,null),u=!0},i(a){u||(M(l.$$.fragment,a),u=!0)},o(a){W(l.$$.fragment,a),u=!1},d(a){a&&m(s),oe(l)}}}function vt(e){let s,t,l,u=e[2].connect.hostLabel()+"",a,p,o,N,$,k,b,d,w,L,B=e[2].connect.sslLabel()+"",U,z,i,T,P,C,K,D,c,R=e[2].connect.idLabel()+"",I,j,E,x,Q,V,h,A,S,X,Y=e[2].connect.passwordLabel()+"",fe,pe,H,Le,Pe,Ae,ce,Ve,de,Z,Te,Ee=e[2].connect.connect()+"",Ce,J,De,je;b=new _e({props:{name:"host",result:e[1]}}),C=new _e({props:{name:"ssl",result:e[1]}}),h=new _e({props:{name:"id",result:e[1]}}),ce=new _e({props:{name:"password",result:e[1]}});let F=e[0]&&Je();return{c(){s=v("form"),t=v("div"),l=v("label"),a=ee(u),p=O(),o=v("input"),k=O(),ne(b.$$.fragment),d=O(),w=v("div"),L=v("label"),U=ee(B),z=O(),i=v("input"),P=O(),ne(C.$$.fragment),K=O(),D=v("div"),c=v("label"),I=ee(R),j=O(),E=v("input"),V=O(),ne(h.$$.fragment),A=O(),S=v("div"),X=v("label"),fe=ee(Y),pe=O(),H=v("input"),Ae=O(),ne(ce.$$.fragment),Ve=O(),de=v("div"),Z=v("button"),F&&F.c(),Te=O(),Ce=ee(Ee),this.h()},l(f){s=g(f,"FORM",{});var _=y(s);t=g(_,"DIV",{class:!0});var ue=y(t);l=g(ue,"LABEL",{for:!0});var ke=y(l);a=te(ke,u),ke.forEach(m),p=q(ue),o=g(ue,"INPUT",{class:!0,type:!0,name:!0,placeholder:!0}),k=q(ue),re(b.$$.fragment,ue),ue.forEach(m),d=q(_),w=g(_,"DIV",{class:!0});var ie=y(w);L=g(ie,"LABEL",{for:!0});var Ie=y(L);U=te(Ie,B),Ie.forEach(m),z=q(ie),i=g(ie,"INPUT",{class:!0,type:!0,name:!0}),P=q(ie),re(C.$$.fragment,ie),ie.forEach(m),K=q(_),D=g(_,"DIV",{class:!0});var he=y(D);c=g(he,"LABEL",{for:!0});var $e=y(c);I=te($e,R),$e.forEach(m),j=q(he),E=g(he,"INPUT",{class:!0,type:!0,name:!0,autocomplete:!0,placeholder:!0}),V=q(he),re(h.$$.fragment,he),he.forEach(m),A=q(_),S=g(_,"DIV",{class:!0});var me=y(S);X=g(me,"LABEL",{for:!0});var He=y(X);fe=te(He,Y),He.forEach(m),pe=q(me),H=g(me,"INPUT",{class:!0,type:!0,name:!0,autocomplete:!0,placeholder:!0}),Ae=q(me),re(ce.$$.fragment,me),me.forEach(m),Ve=q(_),de=g(_,"DIV",{class:!0});var Re=y(de);Z=g(Re,"BUTTON",{type:!0,class:!0});var ye=y(Z);F&&F.l(ye),Te=q(ye),Ce=te(ye,Ee),ye.forEach(m),Re.forEach(m),_.forEach(m),this.h()},h(){r(l,"for","host"),r(o,"class",N="w-full "+e[7]("host")),r(o,"type","text"),r(o,"name","host"),r(o,"placeholder",$=e[2].connect.hostPlaceholder()),r(t,"class","mb-2"),r(L,"for","ssl"),r(i,"class",T=e[7]("ssl")),r(i,"type","checkbox"),r(i,"name","ssl"),r(w,"class","mb-2"),r(c,"for","host"),r(E,"class",x="w-full "+e[7]("id")),r(E,"type","text"),r(E,"name","id"),r(E,"autocomplete","username"),r(E,"placeholder",Q=e[2].connect.idPlaceholder()),r(D,"class","mb-2"),r(X,"for","host"),r(H,"class",Le="w-full "+e[7]("password")),r(H,"type","password"),r(H,"name","password"),r(H,"autocomplete","current-password"),r(H,"placeholder",Pe=e[2].connect.passwordPlaceholder()),r(S,"class","mb-2"),r(Z,"type","submit"),r(Z,"class","btn primary flex flex-shrink"),Z.disabled=e[8],r(de,"class","flex flex-row justify-end")},m(f,_){we(f,s,_),n(s,t),n(t,l),n(l,a),n(t,p),n(t,o),le(o,e[6]),n(t,k),ae(b,t,null),n(s,d),n(s,w),n(w,L),n(L,U),n(w,z),n(w,i),i.checked=e[5],n(w,P),ae(C,w,null),n(s,K),n(s,D),n(D,c),n(c,I),n(D,j),n(D,E),le(E,e[4]),n(D,V),ae(h,D,null),n(s,A),n(s,S),n(S,X),n(X,fe),n(S,pe),n(S,H),le(H,e[3]),n(S,Ae),ae(ce,S,null),n(s,Ve),n(s,de),n(de,Z),F&&F.m(Z,null),n(Z,Te),n(Z,Ce),J=!0,De||(je=[G(o,"input",e[14]),G(o,"input",e[9]),G(i,"change",e[15]),G(i,"input",e[9]),G(E,"input",e[16]),G(E,"input",e[9]),G(H,"input",e[17]),G(H,"input",e[9]),G(s,"submit",Ke(e[10]))],De=!0)},p(f,[_]){(!J||_&4)&&u!==(u=f[2].connect.hostLabel()+"")&&se(a,u),(!J||_&128&&N!==(N="w-full "+f[7]("host")))&&r(o,"class",N),(!J||_&4&&$!==($=f[2].connect.hostPlaceholder()))&&r(o,"placeholder",$),_&64&&o.value!==f[6]&&le(o,f[6]);const ue={};_&2&&(ue.result=f[1]),b.$set(ue),(!J||_&4)&&B!==(B=f[2].connect.sslLabel()+"")&&se(U,B),(!J||_&128&&T!==(T=f[7]("ssl")))&&r(i,"class",T),_&32&&(i.checked=f[5]);const ke={};_&2&&(ke.result=f[1]),C.$set(ke),(!J||_&4)&&R!==(R=f[2].connect.idLabel()+"")&&se(I,R),(!J||_&128&&x!==(x="w-full "+f[7]("id")))&&r(E,"class",x),(!J||_&4&&Q!==(Q=f[2].connect.idPlaceholder()))&&r(E,"placeholder",Q),_&16&&E.value!==f[4]&&le(E,f[4]);const ie={};_&2&&(ie.result=f[1]),h.$set(ie),(!J||_&4)&&Y!==(Y=f[2].connect.passwordLabel()+"")&&se(fe,Y),(!J||_&128&&Le!==(Le="w-full "+f[7]("password")))&&r(H,"class",Le),(!J||_&4&&Pe!==(Pe=f[2].connect.passwordPlaceholder()))&&r(H,"placeholder",Pe),_&8&&H.value!==f[3]&&le(H,f[3]);const Ie={};_&2&&(Ie.result=f[1]),ce.$set(Ie),f[0]?F?_&1&&M(F,1):(F=Je(),F.c(),M(F,1),F.m(Z,Te)):F&&(Oe(),W(F,1,1,()=>{F=null}),qe()),(!J||_&4)&&Ee!==(Ee=f[2].connect.connect()+"")&&se(Ce,Ee),(!J||_&256)&&(Z.disabled=f[8])},i(f){J||(M(b.$$.fragment,f),M(C.$$.fragment,f),M(h.$$.fragment,f),M(ce.$$.fragment,f),M(F),J=!0)},o(f){W(b.$$.fragment,f),W(C.$$.fragment,f),W(h.$$.fragment,f),W(ce.$$.fragment,f),W(F),J=!1},d(f){f&&m(s),oe(b),oe(C),oe(h),oe(ce),F&&F.d(),De=!1,ze(je)}}}const gt=e=>Ye((s={},t)=>{t.length&&(Ze(t),ve("host",e.errors.message.required(),()=>{ge(s.host).isNotBlank()}),ve("ssl",e.errors.message.required(),()=>{ge(s.ssl).isNotBlank()}),ve("id",e.errors.message.required(),()=>{ge(s.id).isNotBlank()}),ve("password",e.errors.message.required(),()=>{ge(s.password).isNotBlank()}))});function bt(e,s,t){let l,u,a,p,o,N,$,k,b,d;be(e,Ue,c=>t(2,b=c)),be(e,ft,c=>t(13,d=c));let{onConnect:w}=s;const L=new Set,B=et(()=>{o({host:l,ssl:u,id:a,password:p},Array.from(L)).done(c=>{t(1,N=c)}),L.clear()},300);function U(){L.add("host"),L.add("ssl"),L.add("id"),L.add("password"),B(),B.flush()}function z(c){c.currentTarget.value=c.currentTarget.value,L.add(c.currentTarget.name),B()}let i=!1;async function T(c){try{if(t(0,i=!0),U(),N.isValid()){const R={host:l,ssl:u,id:a,password:p};await rt(R),pt(R),tt(await at.iceServers()),await w(R)}}catch(R){await xe(R)}finally{t(0,i=!1)}}function P(){l=this.value,t(6,l),t(13,d)}function C(){u=this.checked,t(5,u),t(13,d)}function K(){a=this.value,t(4,a),t(13,d)}function D(){p=this.value,t(3,p),t(13,d)}return e.$$set=c=>{"onConnect"in c&&t(11,w=c.onConnect)},e.$$.update=()=>{e.$$.dirty&8192&&t(6,l=(d==null?void 0:d.host)||"p2p.aicacia.com"),e.$$.dirty&8192&&t(5,u=(d==null?void 0:d.ssl)!==!1),e.$$.dirty&8192&&t(4,a=(d==null?void 0:d.id)||""),e.$$.dirty&8192&&t(3,p=(d==null?void 0:d.password)||""),e.$$.dirty&4&&t(12,o=gt(b)),e.$$.dirty&4096&&t(1,N=o.get()),e.$$.dirty&1&&t(8,$=i),e.$$.dirty&2&&t(7,k=Qe(N,{untested:"untested",tested:"tested",invalid:"invalid",valid:"valid",warning:"warning"}))},[i,N,b,p,a,u,l,k,$,z,T,w,o,d,P,C,K,D]}class wt extends Be{constructor(s){super(),Se(this,s,bt,vt,Ne,{onConnect:11})}}function We(e){let s,t,l,u;return l=new Xe({}),{c(){s=v("div"),t=v("div"),ne(l.$$.fragment),this.h()},l(a){s=g(a,"DIV",{class:!0});var p=y(s);t=g(p,"DIV",{class:!0});var o=y(t);re(l.$$.fragment,o),o.forEach(m),p.forEach(m),this.h()},h(){r(t,"class","inline-block h-6 w-6"),r(s,"class","mr-2 flex flex-row justify-center")},m(a,p){we(a,s,p),n(s,t),ae(l,t,null),u=!0},i(a){u||(M(l.$$.fragment,a),u=!0)},o(a){W(l.$$.fragment,a),u=!1},d(a){a&&m(s),oe(l)}}}function Et(e){let s,t,l,u=e[2].connect.hostLabel()+"",a,p,o,N,$,k,b,d,w,L,B=e[2].connect.sslLabel()+"",U,z,i,T,P,C,K,D,c,R,I=e[2].connect.connect()+"",j,E,x,Q;b=new _e({props:{name:"host",result:e[1]}}),C=new _e({props:{name:"ssl",result:e[1]}});let V=e[0]&&We();return{c(){s=v("form"),t=v("div"),l=v("label"),a=ee(u),p=O(),o=v("input"),k=O(),ne(b.$$.fragment),d=O(),w=v("div"),L=v("label"),U=ee(B),z=O(),i=v("input"),P=O(),ne(C.$$.fragment),K=O(),D=v("div"),c=v("button"),V&&V.c(),R=O(),j=ee(I),this.h()},l(h){s=g(h,"FORM",{});var A=y(s);t=g(A,"DIV",{class:!0});var S=y(t);l=g(S,"LABEL",{for:!0});var X=y(l);a=te(X,u),X.forEach(m),p=q(S),o=g(S,"INPUT",{class:!0,type:!0,name:!0,placeholder:!0}),k=q(S),re(b.$$.fragment,S),S.forEach(m),d=q(A),w=g(A,"DIV",{class:!0});var Y=y(w);L=g(Y,"LABEL",{for:!0});var fe=y(L);U=te(fe,B),fe.forEach(m),z=q(Y),i=g(Y,"INPUT",{class:!0,type:!0,name:!0}),P=q(Y),re(C.$$.fragment,Y),Y.forEach(m),K=q(A),D=g(A,"DIV",{class:!0});var pe=y(D);c=g(pe,"BUTTON",{type:!0,class:!0});var H=y(c);V&&V.l(H),R=q(H),j=te(H,I),H.forEach(m),pe.forEach(m),A.forEach(m),this.h()},h(){r(l,"for","host"),r(o,"class",N="w-full "+e[5]("host")),r(o,"type","text"),r(o,"name","host"),r(o,"placeholder",$=e[2].connect.hostPlaceholder()),r(t,"class","mb-2"),r(L,"for","ssl"),r(i,"class",T=e[5]("ssl")),r(i,"type","checkbox"),r(i,"name","ssl"),r(w,"class","mb-2"),r(c,"type","submit"),r(c,"class","btn primary flex flex-shrink"),c.disabled=e[6],r(D,"class","flex flex-row justify-end")},m(h,A){we(h,s,A),n(s,t),n(t,l),n(l,a),n(t,p),n(t,o),le(o,e[4]),n(t,k),ae(b,t,null),n(s,d),n(s,w),n(w,L),n(L,U),n(w,z),n(w,i),i.checked=e[3],n(w,P),ae(C,w,null),n(s,K),n(s,D),n(D,c),V&&V.m(c,null),n(c,R),n(c,j),E=!0,x||(Q=[G(o,"input",e[12]),G(o,"input",e[7]),G(i,"change",e[13]),G(i,"input",e[7]),G(s,"submit",Ke(e[8]))],x=!0)},p(h,[A]){(!E||A&4)&&u!==(u=h[2].connect.hostLabel()+"")&&se(a,u),(!E||A&32&&N!==(N="w-full "+h[5]("host")))&&r(o,"class",N),(!E||A&4&&$!==($=h[2].connect.hostPlaceholder()))&&r(o,"placeholder",$),A&16&&o.value!==h[4]&&le(o,h[4]);const S={};A&2&&(S.result=h[1]),b.$set(S),(!E||A&4)&&B!==(B=h[2].connect.sslLabel()+"")&&se(U,B),(!E||A&32&&T!==(T=h[5]("ssl")))&&r(i,"class",T),A&8&&(i.checked=h[3]);const X={};A&2&&(X.result=h[1]),C.$set(X),h[0]?V?A&1&&M(V,1):(V=We(),V.c(),M(V,1),V.m(c,R)):V&&(Oe(),W(V,1,1,()=>{V=null}),qe()),(!E||A&4)&&I!==(I=h[2].connect.connect()+"")&&se(j,I),(!E||A&64)&&(c.disabled=h[6])},i(h){E||(M(b.$$.fragment,h),M(C.$$.fragment,h),M(V),E=!0)},o(h){W(b.$$.fragment,h),W(C.$$.fragment,h),W(V),E=!1},d(h){h&&m(s),oe(b),oe(C),V&&V.d(),x=!1,ze(Q)}}}const kt=e=>Ye((s={},t)=>{t.length&&(Ze(t),ve("host",e.errors.message.required(),()=>{ge(s.host).isNotBlank()}),ve("ssl",e.errors.message.required(),()=>{ge(s.ssl).isNotBlank()}))});function It(e,s,t){let l,u,a,p,o,N,$,k;be(e,Ue,P=>t(2,$=P)),be(e,dt,P=>t(11,k=P));let{onConnect:b}=s;const d=new Set,w=et(()=>{a({host:l,ssl:u},Array.from(d)).done(P=>{t(1,p=P)}),d.clear()},300);function L(){d.add("host"),d.add("ssl"),w(),w.flush()}function B(P){P.currentTarget.value=P.currentTarget.value,d.add(P.currentTarget.name),w()}let U=!1;async function z(P){try{if(t(0,U=!0),L(),p.isValid()){const C={host:l,ssl:u},K=new ot({...ut,basePath:it(C.host,C.ssl)}),D=new ct(K);tt(await D.iceServers()),ht(C),await b(C)}}catch(C){await xe(C)}finally{t(0,U=!1)}}function i(){l=this.value,t(4,l),t(11,k)}function T(){u=this.checked,t(3,u),t(11,k)}return e.$$set=P=>{"onConnect"in P&&t(9,b=P.onConnect)},e.$$.update=()=>{e.$$.dirty&2048&&t(4,l=(k==null?void 0:k.host)||"127.0.0.1:3000"),e.$$.dirty&2048&&t(3,u=(k==null?void 0:k.ssl)||!1),e.$$.dirty&4&&t(10,a=kt($)),e.$$.dirty&1024&&t(1,p=a.get()),e.$$.dirty&1&&t(6,o=U),e.$$.dirty&2&&t(5,N=Qe(p,{untested:"untested",tested:"tested",invalid:"invalid",valid:"valid",warning:"warning"}))},[U,p,$,u,l,N,o,B,z,b,a,k,i,T]}class Lt extends Be{constructor(s){super(),Se(this,s,It,Et,Ne,{onConnect:9})}}function Pt(e){let s,t;return s=new Lt({props:{onConnect:e[2]}}),{c(){ne(s.$$.fragment)},l(l){re(s.$$.fragment,l)},m(l,u){ae(s,l,u),t=!0},p:Ge,i(l){t||(M(s.$$.fragment,l),t=!0)},o(l){W(s.$$.fragment,l),t=!1},d(l){oe(s,l)}}}function Tt(e){let s,t;return s=new wt({props:{onConnect:e[2]}}),{c(){ne(s.$$.fragment)},l(l){re(s.$$.fragment,l)},m(l,u){ae(s,l,u),t=!0},p:Ge,i(l){t||(M(s.$$.fragment,l),t=!0)},o(l){W(s.$$.fragment,l),t=!1},d(l){oe(s,l)}}}function Ct(e){let s,t,l,u,a,p,o=e[0].connect.connect()+"",N,$,k,b,d=e[0].connect.httpAccess()+"",w,L,B=e[0].connect.p2pAccess()+"",U,z,i,T,P,C,K;document.title=s=e[0].connect.connect();const D=[Tt,Pt],c=[];function R(I,j){return I[1]==="p2p"?0:I[1]==="http"?1:-1}return~(i=R(e))&&(T=c[i]=D[i](e)),{c(){t=O(),l=v("div"),u=v("div"),a=v("div"),p=v("h1"),N=ee(o),$=O(),k=v("select"),b=v("option"),w=ee(d),L=v("option"),U=ee(B),z=O(),T&&T.c(),this.h()},l(I){lt("svelte-l958id",document.head).forEach(m),t=q(I),l=g(I,"DIV",{class:!0});var E=y(l);u=g(E,"DIV",{class:!0});var x=y(u);a=g(x,"DIV",{class:!0});var Q=y(a);p=g(Q,"H1",{class:!0});var V=y(p);N=te(V,o),V.forEach(m),$=q(Q),k=g(Q,"SELECT",{name:!0});var h=y(k);b=g(h,"OPTION",{});var A=y(b);w=te(A,d),A.forEach(m),L=g(h,"OPTION",{});var S=y(L);U=te(S,B),S.forEach(m),h.forEach(m),z=q(Q),T&&T.l(Q),Q.forEach(m),x.forEach(m),E.forEach(m),this.h()},h(){r(p,"class","mb-1"),b.__value="http",le(b,b.__value),L.__value="p2p",le(L,L.__value),r(k,"name","access"),e[1]===void 0&&st(()=>e[3].call(k)),r(a,"class","flex flex-grow flex-col bg-white p-4 shadow dark:bg-gray-800"),r(u,"class","mx-auto flex w-full flex-shrink flex-col p-4 py-10 md:w-72"),r(l,"class","flex flex-grow flex-col justify-end md:justify-start")},m(I,j){we(I,t,j),we(I,l,j),n(l,u),n(u,a),n(a,p),n(p,N),n(a,$),n(a,k),n(k,b),n(b,w),n(k,L),n(L,U),Fe(k,e[1],!0),n(a,z),~i&&c[i].m(a,null),P=!0,C||(K=G(k,"change",e[3]),C=!0)},p(I,[j]){(!P||j&1)&&s!==(s=I[0].connect.connect())&&(document.title=s),(!P||j&1)&&o!==(o=I[0].connect.connect()+"")&&se(N,o),(!P||j&1)&&d!==(d=I[0].connect.httpAccess()+"")&&se(w,d),(!P||j&1)&&B!==(B=I[0].connect.p2pAccess()+"")&&se(U,B),j&2&&Fe(k,I[1]);let E=i;i=R(I),i===E?~i&&c[i].p(I,j):(T&&(Oe(),W(c[E],1,1,()=>{c[E]=null}),qe()),~i?(T=c[i],T?T.p(I,j):(T=c[i]=D[i](I),T.c()),M(T,1),T.m(a,null)):T=null)},i(I){P||(M(T),P=!0)},o(I){W(T),P=!1},d(I){I&&(m(t),m(l)),~i&&c[i].d(),C=!1,K()}}}function yt(e,s,t){let l,u;be(e,Ue,o=>t(0,l=o)),be(e,Me,o=>t(1,u=o));async function a(){await mt(`${_t}/`)}function p(){u=nt(this),Me.set(u)}return[l,u,a,p]}class jt extends Be{constructor(s){super(),Se(this,s,yt,Ct,Ne,{})}}export{jt as component};
