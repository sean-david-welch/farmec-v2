import{t as m,p as c,j as s,q as d,r as l,u as a,L as x}from"./index-IaYqx9tQ.js";import{s as j}from"./Account.module-Df6Oqq_9.js";import{W as p}from"./WarrantyForm-DW6tux2h.js";import{E as u}from"./Error-C4yaTeAb.js";import{L as h}from"./LoginForm-D3daYyf1.js";import"./faPenToSquare-Cu4cs7mJ.js";import"./Blogs.module-m-DV5lDv.js";const F=()=>{const{isAdmin:e,isAuthenticated:t}=m(),{data:i,isLoading:n,isError:o}=c("warranty");return o?s.jsx(u,{}):n?s.jsx(d,{}):s.jsx("section",{id:"warranty",children:t?s.jsxs(l.Fragment,{children:[s.jsx("h1",{className:a.sectionHeading,children:"Warranty Claims:"}),s.jsx(p,{}),e&&i&&i.map(r=>s.jsxs("div",{className:j.warrantyView,children:[s.jsxs("h1",{className:a.mainHeading,children:[r.dealer," -- ",r.owner_name]}),s.jsx("button",{className:a.btnForm,children:s.jsx(x,{to:`/warranty/${r.id}`,children:"View Claim"})})]},r.id))]}):s.jsx("div",{className:a.loginSection,children:s.jsx(h,{})})})};export{F as default};
