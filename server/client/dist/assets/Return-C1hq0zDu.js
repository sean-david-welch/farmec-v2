import{r as a,c as m,j as s,N as h,u as t}from"./index-yDAfzZuS.js";const f=()=>{const n=m.baseUrl,[o,i]=a.useState(null),[r,c]=a.useState(null);return a.useEffect(()=>{const l=window.location.search,u=new URLSearchParams(l).get("session_id");fetch(`${n}/api/checkout/session-status?session_id=${u}`).then(e=>e.json()).then(e=>{console.log("data",e),i(e.status),c(e.customer_email)})},[n]),o==="open"?s.jsx(h,{to:"/checkout"}):o==="complete"?s.jsx("section",{id:"success",children:s.jsx("div",{className:t.loginSection,children:s.jsxs("div",{className:t.loginForm,children:[s.jsx("h1",{className:t.mainHeading,children:"Payment Complete"}),s.jsxs("p",{className:t.paragraph,children:["We appreciate your business! A confirmation email will be sent to ",r,". If you have any questions, please email"," ",s.jsx("a",{href:"mailto:orders@example.com",children:"info@farmec.ie"}),"."]})]})})}):null};export{f as default};
