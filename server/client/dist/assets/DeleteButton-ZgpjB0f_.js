import{D as d,G as v,j as o,q as m,u as g,F as j}from"./index-CYAi8XUF.js";var u={};(function(e){Object.defineProperty(e,"__esModule",{value:!0});var i="fas",r="trash",n=448,s=512,a=[],t="f1f8",c="M135.2 17.7L128 32H32C14.3 32 0 46.3 0 64S14.3 96 32 96H416c17.7 0 32-14.3 32-32s-14.3-32-32-32H320l-7.2-14.3C307.4 6.8 296.3 0 284.2 0H163.8c-12.1 0-23.2 6.8-28.6 17.7zM416 128H32L53.2 467c1.6 25.3 22.6 45 47.9 45H346.9c25.3 0 46.3-19.7 47.9-45L416 128z";e.definition={prefix:i,iconName:r,icon:[n,s,a,t,c]},e.faTrash=e.definition,e.prefix=i,e.iconName=r,e.width=n,e.height=s,e.ligatures=a,e.unicode=t,e.svgPathData=c,e.aliases=a})(u);const H=({id:e,resourceKey:i,navigateBack:r})=>{const n=d(),{mutateAsync:s,isError:a,error:t,isPending:c}=v(i,e);if(c)return o.jsx(m,{});async function f(l){l.preventDefault();try{r&&n("/"),await s()}catch(h){console.error("Error deleting resource:",h)}}return o.jsxs("button",{className:g.btnForm,onClick:f,children:[o.jsx(j,{icon:u.faTrash}),a&&o.jsxs("p",{children:["Error: ",t.message]})]})};export{H as D};
