import{t as l,y as m,A as u,r as x,j as s,q as j,u as t}from"./index-Bwc28JlK.js";import{s as e}from"./Account.module-Df6Oqq_9.js";import{E as f}from"./Error-BoH_vykE.js";import{R as p}from"./RegistrationForm-CVgElwCb.js";import{D as h}from"./DeleteButton-Cog848pb.js";import{D as g}from"./DownloadPdf-C9abSQ1c.js";import"./faPenToSquare-D69RhHBK.js";const B=()=>{const{isAdmin:n}=l(),i=m().id,{data:r,isLoading:o,isError:d}=u("registrations",i);return x.useEffect(()=>{},[i]),d?s.jsx(f,{}):o?s.jsx(j,{}):r?r&&s.jsxs("section",{id:"warranty-detail",children:[s.jsxs("h1",{className:t.sectionHeading,children:["Machine Registration: ",r==null?void 0:r.dealer_name," - ",r==null?void 0:r.owner_name]}),s.jsxs("div",{className:e.warrantyDetail,children:[Object.entries(r).map(([a,c])=>{if(a!=="id"&&a!=="created"&&a!=="parts")return s.jsxs("div",{className:e.warrantyGrid,children:[s.jsx("div",{className:e.label,children:a}),s.jsx("div",{className:e.value,children:String(c)})]},a)}),n&&r.id&&s.jsxs("div",{className:t.optionsBtn,children:[s.jsx(p,{id:r.id}),s.jsx(h,{id:r.id,resourceKey:"registrations",navigateBack:!0})]})]}),s.jsx(g,{registration:r})]}):s.jsx("section",{id:"warranty-detail",children:s.jsx("div",{children:"Warranty claim not found"})})};export{B as default};
