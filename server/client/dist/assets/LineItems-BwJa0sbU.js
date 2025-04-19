import{r as j,s as u,j as e,q as f,u as t,F as p,p as q}from"./index-C6xlo225.js";import{s as C}from"./Account.module-Df6Oqq_9.js";import{f as g,F as U}from"./faPenToSquare-zr5auVJH.js";import{u as A}from"./aws-eKdh3-Wd.js";import{D as H}from"./DeleteButton-Ct5XCj2J.js";import{E as R}from"./Error-Bfq843Gu.js";const h=r=>[{name:"name",label:"Name",type:"text",placeholder:"Enter name",defaultValue:r==null?void 0:r.name},{name:"price",label:"Price",type:"number",placeholder:"Enter price",defaultValue:r==null?void 0:r.price},{name:"image",label:"Image",type:"file",placeholder:"Upload image"}],x=({id:r,lineItem:n})=>{const[o,s]=j.useState(!1),F=n?h(n):h(),{mutateAsync:b,isError:E,error:y,isPending:L}=u("lineitems"),{mutateAsync:P,isError:N,error:S,isPending:w}=u("lineitems",r),m=r?S:y,d=r?N:E,D=r?P:b;async function T(a){a.preventDefault();const l=new FormData(a.currentTarget),i=l.get("image"),v={name:l.get("name"),price:parseFloat(l.get("price")),image:i?i.name:"null"};try{const c=await D(v);if(i){const V={imageFile:i,presignedUrl:c.presignedUrl};await A(V)}c&&!d&&s(!1)}catch(c){console.error("error creating lineItem",c)}}return L||w?e.jsx(f,{}):e.jsxs("section",{id:"form",children:[e.jsx("button",{className:t.btnForm,onClick:()=>s(!o),children:r?e.jsx(p,{icon:g.faPenToSquare}):e.jsxs("div",{children:["Create Line Item",e.jsx(p,{icon:g.faPenToSquare})]})}),e.jsxs(U,{visible:o,onClose:()=>s(!1),children:[e.jsxs("form",{className:t.form,onSubmit:T,encType:"multipart/form-data",children:[e.jsx("h1",{className:t.mainHeading,children:"LineItem Form"}),F.map(a=>e.jsxs("div",{children:[e.jsx("label",{htmlFor:a.name,children:a.label}),e.jsx("input",{type:a.type,name:a.name,id:a.name,placeholder:a.placeholder,defaultValue:a.defaultValue})]},a.name)),e.jsx("button",{className:t.btnForm,type:"submit",children:"Submit"})]}),d&&e.jsxs("p",{children:["Error: ",m==null?void 0:m.message]})]})]})},J=()=>{const{data:r,isLoading:n,isError:o}=q("lineitems");return o?e.jsx(R,{}):n?e.jsx(f,{}):e.jsx("section",{id:"lineItems",children:e.jsxs(j.Fragment,{children:[e.jsx("h1",{className:t.sectionHeading,children:"Product Line Items:"}),e.jsx(x,{}),r&&r.map(s=>e.jsxs("div",{className:C.productView,children:[e.jsxs("h1",{className:t.mainHeading,children:[s.name," -- ",s.price]}),e.jsx("img",{src:s.image,alt:"line item image",width:400,height:400}),s.id&&e.jsxs("div",{className:t.optionsBtn,children:[e.jsx(x,{id:s.id,lineItem:s}),e.jsx(H,{id:s.id,resourceKey:"lineitems"})]})]},s.id))]})})};export{J as default};
