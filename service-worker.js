const n=location.pathname.split("/").slice(0,-1).join("/"),r=[n+"/internal/immutable/entry/app.C4EQu26R.js",n+"/internal/immutable/nodes/0.QwHJBOED.js",n+"/internal/immutable/assets/0.BMP5Nx9q.css",n+"/internal/immutable/nodes/1.BVzNcXTi.js",n+"/internal/immutable/nodes/2.Cs4SrbXG.js",n+"/internal/immutable/nodes/3.DdjQg_nz.js",n+"/internal/immutable/nodes/4.MnA3K_QC.js",n+"/internal/immutable/nodes/5.BwDFOc5Z.js",n+"/internal/immutable/nodes/6.C0UHfz0i.js",n+"/internal/immutable/nodes/7.BzYWc-AS.js",n+"/internal/immutable/nodes/8.Ci266FKf.js",n+"/internal/immutable/chunks/Dropdown.D45euHi8.js",n+"/internal/immutable/assets/Header.Ckn9IwLX.css",n+"/internal/immutable/chunks/Header.D_zR5CJq.js",n+"/internal/immutable/chunks/InputResults.C9hpT5ea.js",n+"/internal/immutable/chunks/control.CYgJF_JY.js",n+"/internal/immutable/chunks/debounce.DpwXS7NU.js",n+"/internal/immutable/chunks/entry.C7xj7j3Z.js",n+"/internal/immutable/chunks/httpAccess.oZ7xIxns.js",n+"/internal/immutable/chunks/i18n-svelte.DO6JJ3x7.js",n+"/internal/immutable/chunks/index.De_1-CT6.js",n+"/internal/immutable/chunks/index.DfAfitrp.js",n+"/internal/immutable/chunks/index.Dmtj9H_x.js",n+"/internal/immutable/chunks/index.Pahc87b8.js",n+"/internal/immutable/chunks/index.ibyW_jwT.js",n+"/internal/immutable/chunks/layout.dEUG4Mrt.js",n+"/internal/immutable/chunks/notifications.BA952Gbu.js",n+"/internal/immutable/chunks/paths.CDMbSYKj.js",n+"/internal/immutable/chunks/preload-helper.D6kgxu3v.js",n+"/internal/immutable/chunks/scheduler.CgEFyQ-S.js",n+"/internal/immutable/chunks/user.DvHJfosi.js",n+"/internal/immutable/entry/start.Djjv1tpJ.js",n+"/internal/immutable/chunks/index.BJzPQx0s.js"],m=[n+"/favicon.png",n+"/icon.png",n+"/icon256x256.png",n+"/manifest.json",n+"/robots.txt"],o="1721466213766",l=self,c=`cache-${o}`,u=[...r,...m];l.addEventListener("install",e=>{async function s(){await(await caches.open(c)).addAll(u)}e.waitUntil(s())});l.addEventListener("activate",e=>{async function s(){for(const t of await caches.keys())t!==c&&await caches.delete(t)}e.waitUntil(s())});l.addEventListener("fetch",e=>{if(e.request.method==="GET"){const s=(async()=>{const t=new URL(e.request.url),a=await caches.open(c);if(u.includes(t.pathname))return a.match(t.pathname);try{const i=await fetch(e.request);return i.status===200&&!e.request.url.startsWith("chrome-extension://")&&a.put(e.request,i.clone()),i}catch{return a.match(e.request)}})();e.respondWith(s)}});
