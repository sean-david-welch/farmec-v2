import{r as b,s as S,j as e,q as E,u as o,F as w,w as $,L as U,x as L,t as R,y as q,z as G,H as K}from"./index-yDAfzZuS.js";import{s as u}from"./Suppliers.module-BeK9mI9U.js";import{E as O}from"./Error-BMZseeAU.js";import{D as k}from"./DeleteButton-CZEq8dR_.js";import{f as N,F as B}from"./faPenToSquare-BfmRx4nu.js";import{u as Y}from"./aws-eKdh3-Wd.js";import{S as z}from"./SupplierForm-Dl_M7S6e.js";const T=(t,r)=>[{name:"supplier_id",label:"Supplier",type:"select",options:Array.isArray(t)?t.map(n=>({label:n.name,value:n.id})):[],placeholder:"Select supplier",defaultValue:r==null?void 0:r.supplier_id},{name:"web_url",label:"YouTube URL",type:"text",placeholder:"Enter YouTube URL",defaultValue:r==null?void 0:r.web_url}],M=({id:t,video:r,suppliers:l})=>{const[n,i]=b.useState(!1),g=r?T(l,r):T(l),{mutateAsync:f,isError:s,error:d,isPending:h}=S("videos"),{mutateAsync:p,isError:v,error:V,isPending:I}=S("videos",t),x=t?V:d,j=t?v:s,D=t?p:f;async function P(a){a.preventDefault();const c=new FormData(a.currentTarget),m={supplier_id:c.get("supplier_id"),web_url:c.get("web_url")};try{await D(m)&&!j&&i(!1)}catch(y){console.error("error creating video",y)}}return h||I?e.jsx(E,{}):e.jsxs("section",{id:"form",children:[e.jsx("button",{className:o.btnForm,onClick:()=>i(!n),children:t?e.jsx(w,{icon:N.faPenToSquare}):e.jsxs("div",{children:["Create Video",e.jsx(w,{icon:N.faPenToSquare})]})}),e.jsxs(B,{visible:n,onClose:()=>i(!1),children:[e.jsxs("form",{className:o.form,onSubmit:P,encType:"multipart/form-data",children:[e.jsx("h1",{className:o.mainHeading,children:"Videos Form"}),g.map(a=>{var c;return e.jsxs("div",{children:[e.jsx("label",{htmlFor:a.name,children:a.label}),a.type==="select"?e.jsx("select",{name:a.name,id:a.name,children:(c=a.options)==null?void 0:c.map(m=>e.jsx("option",{value:m.value,children:m.label},m.value))}):e.jsx("input",{type:a.type,name:a.name,id:a.name,placeholder:a.placeholder})]},a.name)}),e.jsx("button",{className:o.btnForm,type:"submit",children:"Submit"})]}),j&&e.jsxs("p",{children:["Error: ",x==null?void 0:x.message]})]})]})},J=({videos:t,isAdmin:r,suppliers:l})=>e.jsxs("section",{id:"videos",children:[e.jsx("h1",{className:o.sectionHeading,children:"Videos"}),r&&e.jsx(M,{suppliers:l}),e.jsx("div",{className:u.videoGrid,children:t.map(n=>e.jsxs("div",{className:u.videoCard,id:n.title||"",children:[e.jsx("h1",{className:o.mainHeading,children:n.title}),e.jsx("iframe",{width:"425",height:"315",className:u.video,src:`https://www.youtube.com/embed/${n.video_id}`,allowFullScreen:!0}),r&&n.id&&e.jsxs("div",{className:o.optionsBtn,children:[e.jsx(M,{id:n.id,suppliers:l,video:n}),e.jsx(k,{id:n.id,resourceKey:"videos"})]})]},n.id))})]}),Q="_machineCard_1u2o4_1",W="_machineGrid_1u2o4_6",X="_machineInfo_1u2o4_22",_={machineCard:Q,machineGrid:W,machineInfo:X},C=(t,r)=>[{name:"supplier_id",label:"Supplier",type:"select",options:Array.isArray(t)?t.map(n=>({label:n.name,value:n.id})):[],placeholder:"Select supplier",defaultValue:r==null?void 0:r.supplier_id},{name:"name",label:"Name",type:"text",placeholder:"Enter name",defaultValue:r==null?void 0:r.name},{name:"machine_image",label:"Machine Image",type:"file",placeholder:"Upload machine image"},{name:"description",label:"Description",type:"text",placeholder:"Enter description",defaultValue:r==null?void 0:r.description},{name:"machine_link",label:"Machine Link",type:"text",placeholder:"Enter machine link",defaultValue:r==null?void 0:r.machine_link}],H=({id:t,machine:r,suppliers:l})=>{const[n,i]=b.useState(!1),g=r?C(l,r):C(l),{mutateAsync:f,isError:s,error:d,isPending:h}=S("machines"),{mutateAsync:p,isError:v,error:V,isPending:I}=S("machines",t),x=t?v:s,j=t?V:d,D=t?p:f;async function P(a){a.preventDefault();const c=new FormData(a.currentTarget),m=c.get("machine_image"),y={supplier_id:c.get("supplier_id"),name:c.get("name"),machine_image:m?m.name:"null",description:c.get("description"),machine_link:c.get("machine_link")};try{const F=await D(y);if(m){const A={imageFile:m,presignedUrl:F.presignedUrl};await Y(A)}F&&!x&&i(!1)}catch(F){console.error("Error creating machine",F)}}return h||I?e.jsx(E,{}):e.jsxs("section",{id:"form",children:[e.jsx("button",{className:o.btnForm,onClick:()=>i(!n),children:t?e.jsx(w,{icon:N.faPenToSquare}):e.jsxs("div",{children:["Create Machine",e.jsx(w,{icon:N.faPenToSquare})]})}),e.jsxs(B,{visible:n,onClose:()=>i(!1),children:[e.jsxs("form",{className:o.form,onSubmit:P,encType:"multipart/form-data",children:[e.jsx("h1",{className:o.mainHeading,children:"Machine Form"}),g.map(a=>{var c;return e.jsxs("div",{children:[e.jsx("label",{htmlFor:a.name,children:a.label}),a.type==="select"?e.jsx("select",{name:a.name,id:a.name,children:(c=a.options)==null?void 0:c.map(m=>e.jsx("option",{value:m.value,defaultValue:a.defaultValue,children:m.label},m.value))}):e.jsx("input",{type:a.type,name:a.name,id:a.name,placeholder:a.placeholder,defaultValue:a.defaultValue})]},a.name)}),e.jsx("button",{className:o.btnForm,type:"submit",children:"Submit"})]}),x&&e.jsxs("p",{children:["Error: ",j==null?void 0:j.message]})]})]})},Z=({machines:t,isAdmin:r})=>{const{suppliers:l}=$();if(!l)return e.jsx(E,{});const n=i=>{i.currentTarget.src="/default.jpg"};return e.jsxs("section",{id:"machines",children:[e.jsx("h1",{className:o.sectionHeading,children:"Machinery"}),t.map(i=>e.jsxs(b.Fragment,{children:[e.jsx("div",{className:_.machineCard,id:i.name,children:e.jsxs("div",{className:_.machineGrid,children:[e.jsx("img",{src:i.machine_image,alt:i.name||"Default Image",className:_.machineImage,width:600,height:600,onError:n}),e.jsxs("div",{className:_.machineInfo,children:[e.jsx("h1",{className:o.mainHeading,children:i.name}),e.jsx("p",{className:o.paragraph,children:i.description}),e.jsx("button",{className:o.btn,children:e.jsxs(U,{to:`/machines/${i.id}`,children:["View Products",e.jsx(w,{icon:L.faRightToBracket})]})})]})]})}),r&&i.id&&e.jsxs("div",{className:o.optionsBtn,children:[e.jsx(H,{id:i.id,machine:i,suppliers:l}),e.jsx(k,{id:i.id,resourceKey:"machines"})]})]},i.id)),r&&e.jsx(H,{suppliers:l})]})},oe=()=>{const{isAdmin:t}=R(),{suppliers:r}=$(),l=q().id,n=["suppliers","supplierMachine","videos"],{data:i,isLoading:g,isError:f}=G(l,n);if(b.useEffect(()=>{},[l]),f)return e.jsx(O,{});if(g)return e.jsx(E,{});const[s,d,h]=i;return e.jsxs(e.Fragment,{children:[e.jsxs(K,{children:[e.jsx("title",{children:s?`${s.name} - Farmec Ireland`:"Supplier - Farmec Ireland"}),e.jsx("meta",{name:"description",content:s?s.description:"Browse our Suppliers and learn more about the machines we offer."}),e.jsx("meta",{property:"og:title",content:s?`${s.name} - Farmec Ireland`:"Supplier - Farmec Ireland"}),e.jsx("meta",{property:"og:description",content:s?s.description:"Browse our Suppliers and learn more about the machines we offer."}),e.jsx("meta",{property:"og:image",content:s!=null&&s.marketing_image?s.marketing_image:"https://www.farmec.ie/farmec_images/Suppliers/sip1250.webp"}),e.jsx("meta",{property:"og:url",content:`https://www.farmec.ie/suppliers/${s==null?void 0:s.id}`}),e.jsx("meta",{property:"og:type",content:"website"}),e.jsx("meta",{name:"twitter:card",content:"summary_large_image"}),e.jsx("meta",{name:"twitter:title",content:s?`${s.name} - Farmec Ireland`:"Supplier - Farmec Ireland"}),e.jsx("meta",{name:"twitter:description",content:s?s.description:"Browse our Suppliers and learn more about the machines we offer."}),e.jsx("meta",{name:"twitter:image",content:s!=null&&s.marketing_image?s.marketing_image:"https://www.farmec.ie/farmec_images/Suppliers/sip1250.webp"}),e.jsx("link",{rel:"canonical",href:`https://www.farmec.ie/suppliers/${s==null?void 0:s.id}`})]}),e.jsxs("section",{id:"supplierDetail",children:[s?e.jsxs(b.Fragment,{children:[e.jsxs("div",{className:u.supplierHeading,children:[e.jsx("h1",{className:o.sectionHeading,children:s.name}),t&&s.id&&e.jsxs("div",{className:o.optionsBtn,children:[e.jsx(z,{id:s.id,supplier:s}),e.jsx(k,{id:s.id,resourceKey:"suppliers"})]})]}),d&&e.jsxs("div",{className:o.index,children:[e.jsx("h1",{className:o.indexHeading,children:"Machines"}),d.map(p=>e.jsx("a",{href:`#${p.name}`,children:e.jsx("h1",{className:o.indexItem,children:p.name})},p.name))]}),e.jsxs("div",{className:u.supplierDetail,children:[e.jsx("img",{src:s.marketing_image??"/default.jpg",alt:"/dafault.jpg",className:u.supplierImage,width:750,height:750}),e.jsx("p",{className:u.supplierDescription,children:s.description})]})]}):null,d?e.jsx(Z,{machines:d,isAdmin:t}):null,h?e.jsx(J,{suppliers:r,videos:h,isAdmin:t}):null]})]})};export{oe as default};
