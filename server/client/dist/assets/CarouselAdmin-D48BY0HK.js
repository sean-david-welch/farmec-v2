import{r as f,s as x,j as e,q as F,u as n,F as h,t as U,p as q}from"./index-Bgw3TDtt.js";import{s as H}from"./Account.module-Df6Oqq_9.js";import{E as R}from"./Error-BF5_-chJ.js";import{f as g,F as V}from"./faPenToSquare-vXjp4h8H.js";import{u as B}from"./aws-eKdh3-Wd.js";import{D as I}from"./DeleteButton-D0Z2v94X.js";const j=s=>[{name:"name",label:"Name",type:"text",placeholder:"Enter name",defaultValue:s==null?void 0:s.name},{name:"image",label:"Image",type:"file",placeholder:"Upload image"}],d=({id:s,carousel:o})=>{const[i,t]=f.useState(!1),a=o?j(o):j(),{mutateAsync:b,isError:E,error:y,isPending:C}=x("carousels"),{mutateAsync:N,isError:S,error:P,isPending:w}=x("carousels",s),c=s?P:y,u=s?S:E,D=s?N:b;async function A(r){r.preventDefault();const p=new FormData(r.currentTarget),m=p.get("image"),T={name:p.get("name"),image:m?m.name:"null"};try{const l=await D(T);if(m){const v={imageFile:m,presignedUrl:l.presignedUrl};await B(v)}l&&!u&&t(!1)}catch(l){console.error("error creating carousel",l)}}return C||w?e.jsx(F,{}):e.jsxs("section",{id:"form",children:[e.jsx("button",{className:n.btnForm,onClick:()=>t(!i),children:s?e.jsx(h,{icon:g.faPenToSquare}):e.jsxs("div",{children:["Create Carousel",e.jsx(h,{icon:g.faPenToSquare})]})}),e.jsxs(V,{visible:i,onClose:()=>t(!1),children:[e.jsxs("form",{className:n.form,onSubmit:A,encType:"multipart/form-data",children:[e.jsx("h1",{className:n.mainHeading,children:"Carousel Form"}),a.map(r=>e.jsxs("div",{children:[e.jsx("label",{htmlFor:r.name,children:r.label}),e.jsx("input",{type:r.type,name:r.name,id:r.name,placeholder:r.placeholder,defaultValue:r.defaultValue})]},r.name)),e.jsx("button",{className:n.btnForm,type:"submit",children:"Submit"})]}),u&&e.jsxs("p",{children:["Error: ",c==null?void 0:c.message]})]})]})},J=()=>{const{isAdmin:s}=U(),{data:o,isError:i,isLoading:t}=q("carousels");return i?e.jsx(R,{}):t?e.jsx(F,{}):o?e.jsxs("section",{id:"carousel",children:[e.jsx("h1",{className:n.sectionHeading,children:"Carousels:"}),e.jsx(d,{}),s?e.jsx(f.Fragment,{children:o.map(a=>e.jsxs("div",{className:H.carouselAdmin,children:[e.jsx("h1",{className:n.mainHeading,children:a.name}),e.jsx("img",{src:a.image,alt:"carousel image",width:400}),s&&a.id&&e.jsxs("div",{className:n.optionsBtn,children:[e.jsx(d,{id:a.id,carousel:a})," ",e.jsx(I,{id:a.id,resourceKey:"carousels"})]})]},a.id))}):null]}):e.jsxs("section",{id:"carousel",children:[e.jsx("h1",{children:"No models found"}),s&&e.jsx(d,{})]})};export{J as default};
