import{r as _,s as b,j as e,q as f,u as i,F as j,A as M,t as H,y as R,z as V,H as $,x as A}from"./index-WRsQAKNN.js";import{E as C}from"./Error-YQUkzED6.js";import{s as g}from"./Suppliers.module-BeK9mI9U.js";import{f as y,F as U}from"./faPenToSquare-Dw8WaajV.js";import{u as q}from"./aws-eKdh3-Wd.js";import{D as L}from"./DeleteButton-BC_5WLWN.js";const F=(n,r)=>[{name:"machine_id",label:"Machine",type:"select",options:[{label:n.name,value:n.id}],placeholder:"Select machine",defaultValue:r==null?void 0:r.machine_id},{name:"name",label:"Name",type:"text",placeholder:"Enter name",defaultValue:r==null?void 0:r.name},{name:"product_image",label:"Product Image",type:"file",placeholder:"Upload product image"},{name:"description",label:"Description",type:"text",placeholder:"Enter description",defaultValue:r==null?void 0:r.description},{name:"product_link",label:"Product Link",type:"text",placeholder:"Enter product link",defaultValue:r==null?void 0:r.product_link}],E=({id:n,product:r,machine:m})=>{const[l,d]=_.useState(!1),t=n?F(m,r):F(m),{mutateAsync:a,isError:p,error:u,isPending:P}=b("products"),{mutateAsync:k,isError:N,error:S,isPending:I}=b("products",n),w=n?N:p,x=n?S:u,v=n?k:a;async function D(s){s.preventDefault();const o=new FormData(s.currentTarget),c=o.get("product_image"),T={machine_id:o.get("machine_id"),name:o.get("name"),product_image:c?c.name:"null",description:o.get("description"),product_link:o.get("product_link")};try{const h=await v(T);if(c){const B={imageFile:c,presignedUrl:h.presignedUrl};await q(B)}h&&!w&&d(!1)}catch(h){console.error("error creating product",h)}}return P||I?e.jsx(f,{}):e.jsxs("section",{id:"form",children:[e.jsx("button",{className:i.btnForm,onClick:()=>d(!l),children:n?e.jsx(j,{icon:y.faPenToSquare}):e.jsxs("div",{children:["Create Product",e.jsx(j,{icon:y.faPenToSquare})]})}),e.jsxs(U,{visible:l,onClose:()=>d(!1),children:[e.jsxs("form",{className:i.form,onSubmit:D,encType:"multipart/form-data",children:[e.jsx("h1",{className:i.mainHeading,children:"Product Form"}),t.map(s=>{var o;return e.jsxs("div",{children:[e.jsx("label",{htmlFor:s.name,children:s.label}),s.type==="select"?e.jsx("select",{name:s.name,id:s.name,children:(o=s.options)==null?void 0:o.map(c=>e.jsx("option",{value:c.value,children:c.label},c.value))}):e.jsx("input",{type:s.type,name:s.name,id:s.name,placeholder:s.placeholder,defaultValue:s.defaultValue})]},s.label)}),e.jsx("button",{className:i.btnForm,type:"submit",children:"Submit"})]}),w&&e.jsxs("p",{children:["Error: ",x==null?void 0:x.message]})]})]})},G=({id:n,isAdmin:r,products:m})=>{const{data:l}=M("machines",n);if(!l)return e.jsx(f,{});const d=t=>{t.currentTarget.src="/default.jpg"};return e.jsx("section",{id:"products",children:e.jsx("div",{className:g.productGrid,children:m.map(t=>e.jsxs("div",{className:g.productCard,id:t.name||"",children:[e.jsx("h1",{className:i.mainHeading,children:t.name}),e.jsx("a",{href:t.product_link||"#",target:"_blank",children:e.jsx("img",{src:t.product_image,alt:"/default.jpg",className:g.productImage,width:500,height:500,onError:d})}),e.jsx("p",{className:i.paragraph,children:t.description}),r&&t.id&&e.jsxs("div",{className:i.optionsBtn,children:[e.jsx(E,{id:t.id,product:t,machine:l}),e.jsx(L,{id:t.id,resourceKey:"products"})]})]},t.id))})})},X=()=>{const{isAdmin:n}=H(),r=R().id,m=["machines","products"],{data:l,isLoading:d,isError:t}=V(r,m);if(_.useEffect(()=>{},[r]),t)return e.jsx(C,{});if(d)return e.jsx(f,{});const[a,p]=l;return e.jsxs(e.Fragment,{children:[e.jsxs($,{children:[e.jsx("title",{children:a?`${a.name} - Farmec Ireland`:"Machine - Farmec Ireland"}),e.jsx("meta",{name:"description",content:a?a.description:"Browse our machines and products to learn more about what we offer."}),e.jsx("meta",{property:"og:title",content:a?`${a.name} - Farmec Ireland`:"Machine - Farmec Ireland"}),e.jsx("meta",{property:"og:description",content:a?a.description:"Browse our machines and products to learn more about what we offer."}),e.jsx("meta",{property:"og:image",content:a!=null&&a.marketing_image?a.marketing_image:"https://www.farmec.ie/farmec_images/Suppliers/sip1250.webp"}),e.jsx("meta",{property:"og:url",content:`https://www.farmec.ie/machines/${a==null?void 0:a.id}`}),e.jsx("meta",{property:"og:type",content:"website"}),e.jsx("meta",{name:"twitter:card",content:"summary_large_image"}),e.jsx("meta",{name:"twitter:title",content:a?`${a.name} - Farmec Ireland`:"Machine - Farmec Ireland"}),e.jsx("meta",{name:"twitter:description",content:a?a.description:"Browse our machines and products to learn more about what we offer."}),e.jsx("meta",{name:"twitter:image",content:a!=null&&a.marketing_image?a.marketing_image:"https://www.farmec.ie/farmec_images/Suppliers/sip1250.webp"}),e.jsx("link",{rel:"canonical",href:`https://www.farmec.ie/machines/${a==null?void 0:a.id}`})]}),e.jsxs("section",{id:"machineDetail",children:[e.jsx("h1",{className:i.sectionHeading,children:"Products"}),n&&e.jsx(E,{machine:a}),p&&e.jsxs("div",{className:i.index,children:[e.jsx("h1",{className:i.indexHeading,children:"Products:"}),p.map(u=>e.jsx("a",{href:`#${u.name}`,children:e.jsx("h1",{className:i.indexItem,children:u.name})},u.name)),e.jsx("button",{className:i.btn,children:e.jsxs("a",{href:a.machine_link||"#",target:"_blank",rel:"noopener noreferrer",children:["Supplier Website",e.jsx(j,{icon:A.faRightToBracket})]})})]}),p&&e.jsx(G,{id:r,products:p,isAdmin:n})]})]})};export{X as default};
