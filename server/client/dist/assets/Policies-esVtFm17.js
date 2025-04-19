import{r as A,s as j,j as e,q as S,u as t,F as y,t as I,v as V,H as _}from"./index-CYAi8XUF.js";import{a as T,p as C,s as h}from"./aboutFields-CxfGLbWm.js";import{E as k}from"./Error-D_75nCVZ.js";import{f as g,F as R}from"./faPenToSquare-H85bz21w.js";import{D}from"./DeleteButton-ZgpjB0f_.js";const H=({id:a,term:o})=>{const[i,n]=A.useState(!1),m=o?T(o):T(),{mutateAsync:l,isError:d,error:s,isPending:b}=j("terms"),{mutateAsync:f,isError:w,error:v,isPending:P}=j("terms",a),c=a?v:s,u=a?w:d,F=a?f:l;async function E(r){r.preventDefault();const p=new FormData(r.currentTarget),N={title:p.get("title"),body:p.get("body")};try{await F(N)&&!u&&n(!1)}catch(x){console.error("error creating term",x)}}return b||P?e.jsx(S,{}):e.jsxs("section",{id:"form",children:[e.jsx("button",{className:t.btnForm,onClick:()=>n(!i),children:a?e.jsx(y,{icon:g.faPenToSquare}):e.jsxs("div",{children:["Create Term",e.jsx(y,{icon:g.faPenToSquare})]})}),e.jsxs(R,{visible:i,onClose:()=>n(!1),children:[e.jsxs("form",{className:t.form,onSubmit:E,encType:"multipart/form-data",children:[e.jsx("h1",{className:t.mainHeading,children:"Terms Form"}),m.map(r=>e.jsxs("div",{children:[e.jsx("label",{htmlFor:r.name,children:r.label}),e.jsx("input",{type:r.type,name:r.name,id:r.name,placeholder:r.placeholder,defaultValue:r.defaultValue})]},r.name)),e.jsx("button",{className:t.btnForm,type:"submit",children:"Submit"})]}),u&&e.jsxs("p",{children:["Error: ",c==null?void 0:c.message]})]})]})},q=({id:a,privacy:o})=>{const[i,n]=A.useState(!1),m=o?C(o):C(),{mutateAsync:l,isError:d,error:s,isPending:b}=j("privacys"),{mutateAsync:f,isError:w,error:v,isPending:P}=j("privacys",a),c=a?v:s,u=a?w:d,F=a?f:l;async function E(r){r.preventDefault();const p=new FormData(r.currentTarget),N={title:p.get("title"),body:p.get("body")};try{await F(N)&&!u&&n(!1)}catch(x){console.error("error creating privacy",x)}}return b||P?e.jsx(S,{}):e.jsxs("section",{id:"form",children:[e.jsx("button",{className:t.btnForm,onClick:()=>n(!i),children:a?e.jsx(y,{icon:g.faPenToSquare}):e.jsxs("div",{children:["Create Privacy",e.jsx(y,{icon:g.faPenToSquare})]})}),e.jsxs(R,{visible:i,onClose:()=>n(!1),children:[e.jsxs("form",{className:t.form,onSubmit:E,encType:"multipart/form-data",children:[e.jsx("h1",{className:t.mainHeading,children:"Privacy Form"}),m.map(r=>e.jsxs("div",{children:[e.jsx("label",{htmlFor:r.name,children:r.label}),e.jsx("input",{type:r.type,name:r.name,id:r.name,placeholder:r.placeholder,defaultValue:r.defaultValue})]},r.name)),e.jsx("button",{className:t.btnForm,type:"submit",children:"Submit"})]}),u&&e.jsxs("p",{children:["Error: ",c==null?void 0:c.message]})]})]})},W=()=>{const{isAdmin:a}=I(),o=["terms","privacys"],{data:i,isLoading:n,isError:m}=V(o);if(m)return e.jsx(k,{});if(n)return e.jsx(S,{});const[l,d]=i;return e.jsxs(e.Fragment,{children:[e.jsxs(_,{children:[e.jsx("title",{children:"Policies - Farmec Ireland"}),e.jsx("meta",{name:"description",content:"Read more about our Privacy Policy and how we use your data"}),e.jsx("meta",{property:"og:title",content:"Policies - Farmec Ireland"}),e.jsx("meta",{property:"og:description",content:"Read more about our Privacy Policy and how we use your data."}),e.jsx("meta",{property:"og:image",content:"https://www.farmec.ie/farmec_images/Suppliers/sip1250.webp"}),e.jsx("meta",{property:"og:url",content:"https://www.farmec.ie/policies"}),e.jsx("meta",{property:"og:type",content:"website"}),e.jsx("meta",{name:"twitter:card",content:"summary_large_image"}),e.jsx("meta",{name:"twitter:title",content:"Policies - Farmec Ireland"}),e.jsx("meta",{name:"twitter:description",content:"Read more about our Privacy Policy and how we use your data."}),e.jsx("meta",{name:"twitter:image",content:"https://www.farmec.ie/farmec_images/Suppliers/sip1250.webp"}),e.jsx("link",{rel:"canonical",href:"https://www.farmec.ie/policies"})]}),e.jsxs("section",{id:"policies",children:[e.jsxs("div",{className:h.terms,children:[e.jsx("h1",{className:t.sectionHeading,children:"Terms of Service"}),a&&e.jsx(H,{}),l.map(s=>e.jsxs("div",{className:h.infoCard,children:[e.jsx("h2",{className:t.mainHeading,children:s.title}),e.jsx("p",{className:t.paragraph,children:s.body}),a&&s.id&&e.jsxs("div",{className:t.optionsBtn,children:[e.jsx(H,{id:s.id,term:s}),e.jsx(D,{id:s.id,resourceKey:"terms"})]})]},s.id))]}),e.jsxs("div",{className:h.privacy,children:[e.jsx("h1",{className:t.sectionHeading,children:"Privacy Policy"}),a&&e.jsx(q,{}),d.map(s=>e.jsxs("div",{className:h.infoCard,children:[e.jsx("h2",{className:t.mainHeading,children:s.title}),e.jsx("p",{className:t.paragraph,children:s.body}),a&&s.id&&e.jsxs("div",{className:t.optionsBtn,children:[e.jsx(q,{id:s.id,privacy:s}),e.jsx(D,{id:s.id,resourceKey:"privacys"})]})]},s.id))]})]})]})};export{W as default};
