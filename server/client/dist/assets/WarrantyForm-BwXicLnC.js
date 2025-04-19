import{r as m,s as h,j as t,q as I,u as a,F as x}from"./index-C6xlo225.js";import{f,F as M}from"./faPenToSquare-zr5auVJH.js";const g=e=>[{name:"dealer",label:"Dealer",type:"text",placeholder:"Enter dealer",defaultValue:e==null?void 0:e.dealer},{name:"dealer_contact",label:"Dealer Contact",type:"text",placeholder:"Enter dealer contact",defaultValue:e==null?void 0:e.dealer_contact},{name:"owner_name",label:"Owner Name",type:"text",placeholder:"Enter owner name",defaultValue:e==null?void 0:e.owner_name},{name:"machine_model",label:"Machine Model",type:"text",placeholder:"Enter machine model",defaultValue:e==null?void 0:e.machine_model},{name:"serial_number",label:"Serial Number",type:"text",placeholder:"Enter serial number",defaultValue:e==null?void 0:e.serial_number},{name:"install_date",label:"Install Date",type:"text",placeholder:"Enter install date",defaultValue:e==null?void 0:e.install_date},{name:"failure_date",label:"Failure Date",type:"text",placeholder:"Enter failure date",defaultValue:e==null?void 0:e.failure_date},{name:"repair_date",label:"Repair Date",type:"text",placeholder:"Enter repair date",defaultValue:e==null?void 0:e.repair_date},{name:"failure_details",label:"Failure Details",type:"text",placeholder:"Enter failure details",defaultValue:e==null?void 0:e.failure_details},{name:"repair_details",label:"Repair Details",type:"text",placeholder:"Enter repair details",defaultValue:e==null?void 0:e.repair_details},{name:"labour_hours",label:"Labour Hours",type:"text",placeholder:"Enter labour hours",defaultValue:e==null?void 0:e.labour_hours},{name:"completed_by",label:"Completed By",type:"text",placeholder:"Enter completed by",defaultValue:e==null?void 0:e.completed_by}],k=(e={},r)=>[{name:`part_number_${r}`,label:"Part Number",type:"text",placeholder:"Enter part number",defaultValue:e==null?void 0:e.part_number},{name:`quantity_needed_${r}`,label:"Quantity Needed",type:"text",placeholder:"Enter quantity needed",defaultValue:e==null?void 0:e.quantity_needed},{name:`invoice_number_${r}`,label:"Invoice Number",type:"text",placeholder:"Enter invoice number",defaultValue:e==null?void 0:e.invoice_number},{name:`part_description_${r}`,label:"Part Description",type:"text",placeholder:"Enter  description",defaultValue:e==null?void 0:e.description}],B=({id:e,warranty:r})=>{const[b,s]=m.useState(!1),[u,E]=m.useState([{part_number:"",quantity_needed:"",invoice_number:"",description:""}]),j=()=>{E([...u,{part_number:"",quantity_needed:"",invoice_number:"",description:""}])},F=r?g(r):g(),{mutateAsync:V,isError:v,error:P,isPending:D}=h("warranty"),{mutateAsync:N,isError:S,error:T,isPending:W}=h("warranty",e),d=e?T:P,i=e?S:v,$=e?N:V;async function A(l){l.preventDefault();const o=new FormData(l.currentTarget),p=u.map((_,c)=>({part_number:l.currentTarget[`part_number_${c}`].value,quantity_needed:l.currentTarget[`quantity_needed_${c}`].value,invoice_number:l.currentTarget[`invoice_number_${c}`].value,description:l.currentTarget[`part_description_${c}`].value})),n={warranty:{dealer:o.get("dealer"),dealer_contact:o.get("dealer_contact"),owner_name:o.get("owner_name"),machine_model:o.get("machine_model"),serial_number:o.get("serial_number"),install_date:o.get("install_date"),failure_date:o.get("failure_date"),repair_date:o.get("repair_date"),failure_details:o.get("failure_details"),repair_details:o.get("repair_details"),labour_hours:o.get("labour_hours"),completed_by:o.get("completed_by")},parts:p};try{await $(n)&&!i&&s(!1)}catch(_){console.error("Failed to create wwarranty claim",_)}s(!1)}return D||W?t.jsx(I,{}):t.jsxs("section",{id:"form",children:[t.jsx("button",{className:a.btnForm,onClick:()=>s(!b),children:e?t.jsx(x,{icon:f.faPenToSquare}):t.jsxs("div",{children:["Warranty Claim",t.jsx(x,{icon:f.faPenToSquare})]})}),t.jsxs(M,{visible:b,onClose:()=>s(!1),children:[t.jsxs("form",{className:a.form,onSubmit:A,children:[t.jsx("h1",{className:a.mainHeading,children:"Warranty Claim Form"}),F.map(l=>t.jsxs("div",{children:[t.jsx("label",{htmlFor:l.name,children:l.label}),t.jsx("input",{type:l.type,name:l.name,id:l.name,placeholder:l.placeholder})]},l.name)),u.map((l,o)=>{const p=k(l,o);return t.jsx("div",{children:p.map(n=>t.jsxs("div",{children:[t.jsx("label",{htmlFor:n.name,children:n.label}),t.jsx("input",{type:n.type,name:n.name,id:n.name,placeholder:n.placeholder})]},n.name))},o)}),t.jsx("button",{type:"button",className:a.btnForm,onClick:j,children:"Add Part"}),t.jsx("button",{className:a.btnForm,type:"submit",children:"Submit"})]}),i&&t.jsxs("p",{children:["Error: ",d==null?void 0:d.message]})]})]})};export{B as W};
