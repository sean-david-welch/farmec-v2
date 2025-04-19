import{r as R,s as u,j as a,q as D,u as i,F as f}from"./index-Bgw3TDtt.js";import{f as _,F as P}from"./faPenToSquare-vXjp4h8H.js";import{u as h}from"./aws-eKdh3-Wd.js";const x=e=>[{name:"name",label:"Name",type:"text",placeholder:"Enter name",defaultValue:e==null?void 0:e.name},{name:"logo_image",label:"Logo Image",type:"file",placeholder:"Upload logo image"},{name:"marketing_image",label:"Marketing Image",type:"file",placeholder:"Upload marketing image"},{name:"description",label:"Description",type:"text",placeholder:"Enter description",defaultValue:e==null?void 0:e.description},{name:"social_facebook",label:"Facebook",type:"text",placeholder:"Enter Facebook URL",defaultValue:e==null?void 0:e.social_facebook},{name:"social_instagram",label:"Instagram",type:"text",placeholder:"Enter Instagram URL",defaultValue:e==null?void 0:e.social_instagram},{name:"social_twitter",label:"Twitter",type:"text",placeholder:"Enter Twitter URL",defaultValue:e==null?void 0:e.social_twitter},{name:"social_linkedin",label:"LinkedIn",type:"text",placeholder:"Enter LinkedIn URL",defaultValue:e==null?void 0:e.social_linkedin},{name:"social_youtube",label:"YouTube",type:"text",placeholder:"Enter YouTube URL",defaultValue:e==null?void 0:e.social_youtube},{name:"social_website",label:"Website",type:"text",placeholder:"Enter website URL",defaultValue:e==null?void 0:e.social_website}],v=({id:e,supplier:g})=>{const[d,s]=R.useState(!1),y=g?x(g):x(),{mutateAsync:F,isError:k,error:E,isPending:w}=u("suppliers"),{mutateAsync:j,isError:S,error:U,isPending:L}=u("suppliers",e),c=e?U:E,b=e?S:k,T=e?j:F;async function V(t){t.preventDefault();const o=new FormData(t.currentTarget),r=o.get("logo_image"),l=o.get("marketing_image"),I={name:o.get("name"),description:o.get("description"),logo_image:r?r.name:"null",marketing_image:l?l.name:"null",social_facebook:o.get("social_facebook"),social_twitter:o.get("social_twitter"),social_instagram:o.get("social_instagram"),social_youtube:o.get("social_youtube"),social_linkedin:o.get("social_linkedin"),social_website:o.get("social_website")};try{const n=await T(I);if(r){const m={imageFile:r,presignedUrl:n.presignedLogoUrl};await h(m)}if(l){const m={imageFile:l,presignedUrl:n.presignedMarketingUrl};await h(m)}n&&!b&&s(!1)}catch(n){console.error("Error creating supplier:",n)}}return w||L?a.jsx(D,{}):a.jsxs("section",{id:"form",children:[a.jsx("button",{className:i.btnForm,onClick:()=>s(!d),children:e?a.jsx(f,{icon:_.faPenToSquare}):a.jsxs("div",{children:["Create Supplier",a.jsx(f,{icon:_.faPenToSquare})]})}),a.jsxs(P,{visible:d,onClose:()=>s(!1),children:[a.jsxs("form",{className:i.form,onSubmit:V,encType:"multipart/form-data",children:[a.jsx("h1",{className:i.mainHeading,children:"Supplier Form"}),y.map(t=>a.jsxs("div",{children:[a.jsx("label",{htmlFor:t.name,children:t.label}),a.jsx("input",{type:t.type,name:t.name,id:t.name,placeholder:t.placeholder,defaultValue:t.defaultValue})]},t.name)),a.jsx("button",{className:i.btnForm,type:"submit",children:"Submit"})]}),b&&a.jsxs("p",{children:["Error: ",c==null?void 0:c.message]})]})]})};export{v as S};
