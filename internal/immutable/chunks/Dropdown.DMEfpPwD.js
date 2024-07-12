import{t as C,n as B,a as D,v as U,w as q,u as z,g as T,b as S,i as V,r as A,o as G,l as W,x as J}from"./scheduler.Bu9aYH8f.js";import{S as F,i as H,e as P,a as L,b as O,h as y,j as _,J as h,k as j,K,F as I,o as k,p as E,L as v,s as Q,c as X,f as Y,d as Z,l as M,m as x,M as $,G as ee,q as te}from"./index.Bw_3EgnJ.js";import{d as oe}from"./debounce.BzpYRo3t.js";function ne(n,t={enabled:!0}){let{enabled:o,eventType:f,nodeForEvent:s,options:a,capture:c}=N(t);function p(r){r.target&&n&&!n.contains(r.target)&&!r.defaultPrevented&&n.dispatchEvent(new CustomEvent("clickoutside",{detail:r}))}return t.enabled!==!1&&s.addEventListener(f,p,a),{update(r){s.removeEventListener(f,p,c),{enabled:o,eventType:f,nodeForEvent:s,options:a,capture:c}=N(r),o&&s.addEventListener(f,p,a)},destroy(){s.removeEventListener(f,p,c)}}}function N(n={}){var t,o;return{enabled:n.enabled??!0,nodeForEvent:((t=n.limit)==null?void 0:t.parent)??document,eventType:n.event??"click",options:n.options,capture:typeof n.options=="object"?(o=n.options)==null?void 0:o.capture:n.options}}function se(n,t="body"){let o;async function f(a){if(t=a,typeof t=="string"){if(o=document.querySelector(t),o===null&&(await C(),o=document.querySelector(t)),o===null)throw new Error(`No element found matching css selector: "${t}"`)}else if(t instanceof HTMLElement)o=t;else throw new TypeError(`Unknown portal target type: ${t===null?"null":typeof t}. Allowed types: string (CSS selector) or HTMLElement.`);o.appendChild(n),n.hidden=!1}function s(){n.parentNode&&n.parentNode.removeChild(n)}return f(t),{update:f,destroy:s}}function ie(n){let t,o,f,s,a,c;const p=n[9].default,r=D(p,n,n[8],null);return{c(){t=P("div"),r&&r.c(),this.h()},l(u){t=L(u,"DIV",{class:!0,role:!0,tabindex:!0});var e=O(t);r&&r.l(e),e.forEach(y),this.h()},h(){_(t,"class","absolute flex max-h-full max-w-full flex-col border-gray-300 bg-gray-50 shadow-md transition-transform duration-75 focus:outline-none dark:border-gray-600 dark:bg-gray-700"),_(t,"role","menu"),_(t,"tabindex","-1"),U(()=>n[11].call(t)),h(t,"scale-0",!n[0]),h(t,"scale-100",n[0]),h(t,"origin-top-right",n[1]==="top-right"),h(t,"origin-top-left",n[1]==="top-left"),h(t,"origin-top",n[1]==="top-center"),h(t,"origin-bottom-right",n[1]==="bottom-right"),h(t,"origin-bottom-left",n[1]==="bottom-left"),h(t,"origin-bottom",n[1]==="bottom-center")},m(u,e){j(u,t,e),r&&r.m(t,null),n[10](t),o=K(t,n[11].bind(t)),s=!0,a||(c=[q(se.call(null,t)),q(f=ne.call(null,t,{event:"pointerdown",enabled:n[2]})),I(t,"clickoutside",n[5])],a=!0)},p(u,[e]){r&&r.p&&(!s||e&256)&&z(r,p,u,u[8],s?S(p,u[8],e,null):T(u[8]),null),f&&V(f.update)&&e&4&&f.update.call(null,{event:"pointerdown",enabled:u[2]}),(!s||e&1)&&h(t,"scale-0",!u[0]),(!s||e&1)&&h(t,"scale-100",u[0]),(!s||e&2)&&h(t,"origin-top-right",u[1]==="top-right"),(!s||e&2)&&h(t,"origin-top-left",u[1]==="top-left"),(!s||e&2)&&h(t,"origin-top",u[1]==="top-center"),(!s||e&2)&&h(t,"origin-bottom-right",u[1]==="bottom-right"),(!s||e&2)&&h(t,"origin-bottom-left",u[1]==="bottom-left"),(!s||e&2)&&h(t,"origin-bottom",u[1]==="bottom-center")},i(u){s||(k(r,u),s=!0)},o(u){E(r,u),s=!1},d(u){u&&y(t),r&&r.d(u),n[10](null),o(),a=!1,A(c)}}}function le(n,t,o){let{$$slots:f={},$$scope:s}=t,{anchor:a}=t,{anchorPosition:c="bottom-right"}=t,{position:p="top-right"}=t,{open:r=!1}=t,{closeOnClickOutside:u=!0}=t,e,b;async function g(){const l=a.getBoundingClientRect();switch(o(3,e.style.top=o(3,e.style.left=o(3,e.style.bottom=o(3,e.style.right="",e),e),e),e),await C(),c){case"top-left":{o(3,e.style.top=`${l.top}px`,e),o(3,e.style.left=`${l.left}px`,e);break}case"top-right":{o(3,e.style.top=`${l.top}px`,e),o(3,e.style.right=`${a.ownerDocument.body.offsetWidth-l.right}px`,e);break}case"top-center":{o(3,e.style.top=`${l.top}px`,e),o(3,e.style.left=`${l.left+l.width*.5-b*.5}px`,e);break}case"bottom-right":{o(3,e.style.top=`${l.bottom}px`,e),o(3,e.style.right=`${a.ownerDocument.body.offsetWidth-l.right}px`,e);break}case"bottom-left":{o(3,e.style.top=`${l.bottom}px`,e),o(3,e.style.left=`${l.left}px`,e);break}case"bottom-center":{o(3,e.style.top=`${l.bottom}px`,e),o(3,e.style.left=`${l.left+l.width*.5-b*.5}px`,e);break}}await C();const w=e.getBoundingClientRect();w.left<0?(o(3,e.style.right="",e),o(3,e.style.left="0px",e)):w.right>a.ownerDocument.body.offsetWidth&&(o(3,e.style.right="0px",e),o(3,e.style.left="",e)),w.top<0?o(3,e.style.top="0px",e):w.top>a.ownerDocument.body.offsetHeight&&(o(3,e.style.top="",e),o(3,e.style.bottom="0px",e))}G(()=>{const l=oe(g,0);return window.addEventListener("resize",l),document.addEventListener("scroll",l,!0),()=>{window.removeEventListener("resize",l),document.removeEventListener("scroll",l,!0)}});function i(l){l.stopPropagation(),r&&o(0,r=!1)}function d(l){W[l?"unshift":"push"](()=>{e=l,o(3,e)})}function m(){b=this.offsetWidth,o(4,b)}return n.$$set=l=>{"anchor"in l&&o(6,a=l.anchor),"anchorPosition"in l&&o(7,c=l.anchorPosition),"position"in l&&o(1,p=l.position),"open"in l&&o(0,r=l.open),"closeOnClickOutside"in l&&o(2,u=l.closeOnClickOutside),"$$scope"in l&&o(8,s=l.$$scope)},n.$$.update=()=>{n.$$.dirty&73&&a&&e&&r&&g()},[r,p,u,e,b,i,a,c,s,f,d,m]}class re extends F{constructor(t){super(),H(this,t,le,ie,B,{anchor:6,anchorPosition:7,position:1,open:0,closeOnClickOutside:2})}}const fe=n=>({}),R=n=>({});function ue(n){let t;const o=n[6].default,f=D(o,n,n[9],null);return{c(){f&&f.c()},l(s){f&&f.l(s)},m(s,a){f&&f.m(s,a),t=!0},p(s,a){f&&f.p&&(!t||a&512)&&z(f,o,s,s[9],t?S(o,s[9],a,null):T(s[9]),null)},i(s){t||(k(f,s),t=!0)},o(s){E(f,s),t=!1},d(s){f&&f.d(s)}}}function ae(n){let t,o,f,s,a,c,p,r;const u=n[6].button,e=D(u,n,n[9],R);function b(i){n[7](i)}let g={anchor:n[4],anchorPosition:n[2],position:n[3],$$slots:{default:[ue]},$$scope:{ctx:n}};return n[0]!==void 0&&(g.open=n[0]),s=new re({props:g}),W.push(()=>v(s,"open",b)),{c(){t=P("div"),o=P("button"),e&&e.c(),f=Q(),X(s.$$.fragment),this.h()},l(i){t=L(i,"DIV",{class:!0});var d=O(t);o=L(d,"BUTTON",{type:!0,"aria-label":!0,class:!0});var m=O(o);e&&e.l(m),m.forEach(y),f=Y(d),Z(s.$$.fragment,d),d.forEach(y),this.h()},h(){_(o,"type","button"),_(o,"aria-label",n[1]),_(o,"class","btn icon primary min-w-fit flex-grow"),_(t,"class","static flex min-w-fit flex-col")},m(i,d){j(i,t,d),M(t,o),e&&e.m(o,null),M(t,f),x(s,t,null),n[8](t),c=!0,p||(r=I(o,"pointerdown",$(ee(n[5]))),p=!0)},p(i,[d]){e&&e.p&&(!c||d&512)&&z(e,u,i,i[9],c?S(u,i[9],d,fe):T(i[9]),R),(!c||d&2)&&_(o,"aria-label",i[1]);const m={};d&16&&(m.anchor=i[4]),d&4&&(m.anchorPosition=i[2]),d&8&&(m.position=i[3]),d&512&&(m.$$scope={dirty:d,ctx:i}),!a&&d&1&&(a=!0,m.open=i[0],J(()=>a=!1)),s.$set(m)},i(i){c||(k(e,i),k(s.$$.fragment,i),c=!0)},o(i){E(e,i),E(s.$$.fragment,i),c=!1},d(i){i&&y(t),e&&e.d(i),te(s),n[8](null),p=!1,r()}}}function ce(n,t,o){let{$$slots:f={},$$scope:s}=t,{name:a=""}=t,{anchorPosition:c=void 0}=t,{position:p=void 0}=t,{open:r=!1}=t;function u(){o(0,r=!r)}let e;function b(i){r=i,o(0,r)}function g(i){W[i?"unshift":"push"](()=>{e=i,o(4,e)})}return n.$$set=i=>{"name"in i&&o(1,a=i.name),"anchorPosition"in i&&o(2,c=i.anchorPosition),"position"in i&&o(3,p=i.position),"open"in i&&o(0,r=i.open),"$$scope"in i&&o(9,s=i.$$scope)},[r,a,c,p,e,u,f,b,g,s]}class me extends F{constructor(t){super(),H(this,t,ce,ae,B,{name:1,anchorPosition:2,position:3,open:0})}}export{me as D,se as p};
