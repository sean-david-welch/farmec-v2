import{r as q,s as g,j as e,q as C,u as n,F as d}from"./index-BdVOQLH1.js";import{f as p,F as N}from"./faPenToSquare-D6SjMlxH.js";import{u as v}from"./aws-eKdh3-Wd.js";import{b}from"./blogFields-Cn3ZjptV.js";const _=({id:s,blog:l})=>{const[c,i]=q.useState(!1),h=l?b(l):b(),{mutateAsync:x,isError:F,error:f,isPending:j}=g("blogs"),{mutateAsync:y,isError:E,error:S,isPending:B}=g("blogs",s),m=s?S:f,u=s?E:F,P=s?y:x;async function w(r){r.preventDefault();const o=new FormData(r.currentTarget),a=o.get("main_image"),T={title:o.get("title"),date:o.get("date"),main_image:a?a.name:"null",subheading:o.get("subheading"),body:o.get("body")};try{const t=await P(T);if(a){const D={imageFile:a,presignedUrl:t.presignedUrl};await v(D)}t&&!u&&i(!1)}catch(t){console.error("error creating blog",t)}}return j||B?e.jsx(C,{}):e.jsxs("section",{id:"form",children:[e.jsx("button",{className:n.btnForm,onClick:()=>i(!c),children:s?e.jsx(d,{icon:p.faPenToSquare}):e.jsxs("div",{children:["Create Blog",e.jsx(d,{icon:p.faPenToSquare})]})}),e.jsxs(N,{visible:c,onClose:()=>i(!1),children:[e.jsxs("form",{className:n.form,onSubmit:w,encType:"multipart/form-data",children:[e.jsx("h1",{className:n.mainHeading,children:"Blog Form"}),h.map(r=>e.jsxs("div",{children:[e.jsx("label",{htmlFor:r.name,children:r.label}),e.jsx("input",{type:r.type,name:r.name,id:r.name,placeholder:r.placeholder,defaultValue:r.defaultValue})]},r.name)),e.jsx("button",{className:n.btnForm,type:"submit",children:"Submit"})]}),u&&e.jsxs("p",{children:["Error: ",m==null?void 0:m.message]})]})]})};export{_ as B};
