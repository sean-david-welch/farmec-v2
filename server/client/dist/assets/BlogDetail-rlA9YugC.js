import{y as d,t as p,A as g,r as i,j as e,q as x,H as j,u as s}from"./index-WRsQAKNN.js";import{s as a}from"./Blogs.module-m-DV5lDv.js";import{E as h}from"./Error-YQUkzED6.js";import{B as u}from"./BlogForm-BYJC_PED.js";import{D as w}from"./DeleteButton-BC_5WLWN.js";import"./faPenToSquare-Dw8WaajV.js";import"./aws-eKdh3-Wd.js";import"./blogFields-Cn3ZjptV.js";const $=()=>{const r=d().id,{isAdmin:n}=p(),{data:t,isLoading:o,isError:m}=g("blogs",r);if(i.useEffect(()=>{},[r]),m)return e.jsx(h,{});if(o)return e.jsx(x,{});const c=l=>{l.currentTarget.src="/default.jpg"};return e.jsxs(e.Fragment,{children:[t&&e.jsxs(j,{children:[e.jsx("title",{children:`${t.title} - Farmec Blog`}),e.jsx("meta",{name:"description",content:t.subheading}),e.jsx("meta",{property:"og:title",content:`${t.title} - Farmec Blog`}),e.jsx("meta",{property:"og:description",content:t.subheading}),e.jsx("meta",{property:"og:image",content:"https://www.farmec.ie/farmec_images/Suppliers/sip1250.webp"}),e.jsx("meta",{property:"og:url",content:`https://www.farmec.ie/blogs/${t.id}`}),e.jsx("meta",{property:"og:type",content:"article"}),e.jsx("meta",{name:"twitter:card",content:"summary_large_image"}),e.jsx("meta",{name:"twitter:title",content:`${t.title} - Farmec Blog`}),e.jsx("meta",{name:"twitter:description",content:t.subheading}),e.jsx("meta",{name:"twitter:image",content:"https://www.farmec.ie/farmec_images/Suppliers/sip1250.webp"}),e.jsx("link",{rel:"canonical",href:`https://www.farmec.ie/blogs/${t.id}`})]}),e.jsx("section",{id:"blog",children:e.jsxs(i.Fragment,{children:[t&&e.jsxs("div",{className:a.blogDetail,children:[e.jsx("h1",{className:s.sectionHeading,children:t.title}),e.jsxs("div",{className:a.blogBody,children:[e.jsx("img",{src:t.main_image,alt:"Blog image",width:600,height:600,onError:c}),e.jsx("h1",{className:s.mainHeading,children:t.subheading}),e.jsx("p",{className:s.paragraph,children:t.body})]})]}),n&&(t==null?void 0:t.id)&&e.jsxs("div",{className:s.optionsBtn,children:[e.jsx(u,{id:t==null?void 0:t.id,blog:t}),e.jsx(w,{id:t==null?void 0:t.id,resourceKey:"blogs",navigateBack:!0})]})]})})]})};export{$ as default};
