import{t as l,p as d,j as e,q as p,H as h,u as s,L as x,F as g,x as j}from"./index-Bwc28JlK.js";import{s as r}from"./Blogs.module-m-DV5lDv.js";import{E as u}from"./Error-BoH_vykE.js";import{B as o}from"./BlogForm-DQUN4wyD.js";import{D as w}from"./DeleteButton-Cog848pb.js";import"./faPenToSquare-D69RhHBK.js";import"./aws-eKdh3-Wd.js";import"./blogFields-Cn3ZjptV.js";const E=()=>{const{isAdmin:i}=l(),{data:a,isLoading:n,isError:m}=d("blogs");if(m)return e.jsx(u,{});if(n)return e.jsx(p,{});const c=t=>{t.currentTarget.src="/default.jpg"};return e.jsxs(e.Fragment,{children:[e.jsxs(h,{children:[e.jsx("title",{children:"Latest Blog Posts - Farmec Ireland"}),e.jsx("meta",{name:"description",content:"Check out the latest blog posts from Farmec Ireland. Stay up to date with our latest news and insights."}),e.jsx("meta",{property:"og:title",content:"Latest Blog Posts - Farmec Ireland"}),e.jsx("meta",{property:"og:description",content:"Check out the latest blog posts from Farmec Ireland. Stay up to date with our latest news and insights."}),e.jsx("meta",{property:"og:image",content:"https://www.farmec.ie/farmec_images/Suppliers/sip1250.webp"}),e.jsx("meta",{property:"og:url",content:"https://www.farmec.ie/blogs"}),e.jsx("meta",{property:"og:type",content:"website"}),e.jsx("meta",{name:"twitter:card",content:"summary_large_image"}),e.jsx("meta",{name:"twitter:title",content:"Latest Blog Posts - Farmec Ireland"}),e.jsx("meta",{name:"twitter:description",content:"Check out the latest blog posts from Farmec Ireland. Stay up to date with our latest news and insights."}),e.jsx("meta",{name:"twitter:image",content:"https://www.farmec.ie/farmec_images/Suppliers/sip1250.webp"}),e.jsx("link",{rel:"canonical",href:"https://www.farmec.ie/blogs"})]}),e.jsxs("section",{id:"blog",children:[e.jsx("h1",{className:s.sectionHeading,children:"Check out our Latest Blog Posts"}),e.jsx("p",{className:s.subHeading,children:"Read our latest news"}),i&&e.jsx(o,{}),a&&e.jsxs("div",{className:s.index,children:[e.jsx("h1",{className:s.indexHeading,children:"Blogs"}),a.map(t=>e.jsx("a",{href:`#${t.title}`,children:e.jsx("h1",{className:s.indexItem,children:t.title})},t.title))]}),a==null?void 0:a.map(t=>e.jsxs("div",{className:r.blogGrid,id:t.title||"",children:[e.jsxs("div",{className:r.blogCard,children:[e.jsx("img",{className:r.blogImage,src:t.main_image,alt:"Blog image",width:300,height:300,onError:c}),e.jsxs("div",{className:r.blogLink,children:[e.jsx("h1",{className:s.mainHeading,children:t.title}),e.jsx("p",{className:s.paragraph,children:t.subheading}),e.jsx("p",{className:s.paragraph,children:t.body}),e.jsx("button",{className:s.btnForm,children:e.jsxs(x,{to:`/blogs/${t.id}`,children:["Read More",e.jsx(g,{icon:j.faRightToBracket})]})})]})]}),i&&t.id&&e.jsxs("div",{className:s.optionsBtn,children:[e.jsx(o,{id:t.id,blog:t}),e.jsx(w,{id:t.id,resourceKey:"blogs"})]})]},t.id))]})]})};export{E as default};
