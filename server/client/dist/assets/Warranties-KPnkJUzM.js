import{t as m,p as c,j as s,q as d,r as l,u as a,L as x}from"./index-CYAi8XUF.js";import{s as j}from"./Account.module-rvB3Z4_I.js";import{W as p}from"./WarrantyForm-C9WBE6Aq.js";import{E as u}from"./Error-D_75nCVZ.js";import{L as h}from"./LoginForm-ChWUsCgP.js";import"./faPenToSquare-H85bz21w.js";import"./Blogs.module-BNC7JxHQ.js";const F=()=>{const{isAdmin:e,isAuthenticated:t}=m(),{data:i,isLoading:n,isError:o}=c("warranty");return o?s.jsx(u,{}):n?s.jsx(d,{}):s.jsx("section",{id:"warranty",children:t?s.jsxs(l.Fragment,{children:[s.jsx("h1",{className:a.sectionHeading,children:"Warranty Claims:"}),s.jsx(p,{}),e&&i&&i.map(r=>s.jsxs("div",{className:j.warrantyView,children:[s.jsxs("h1",{className:a.mainHeading,children:[r.dealer," -- ",r.owner_name]}),s.jsx("button",{className:a.btnForm,children:s.jsx(x,{to:`/warranty/${r.id}`,children:"View Claim"})})]},r.id))]}):s.jsx("div",{className:a.loginSection,children:s.jsx(h,{})})})};export{F as default};
