import{P as s,R as u,r as G,y as Re,c as oe,j as R,u as J}from"./index-CYAi8XUF.js";var ae="https://js.stripe.com/v3",Oe=/^https:\/\/js\.stripe\.com\/v3\/?(\?.*)?$/,X="loadStripe.setLoadParameters was called but an existing Stripe.js script already exists in the document; existing script parameters will be used",Ae=function(){for(var e=document.querySelectorAll('script[src^="'.concat(ae,'"]')),t=0;t<e.length;t++){var n=e[t];if(Oe.test(n.src))return n}return null},z=function(e){var t="",n=document.createElement("script");n.src="".concat(ae).concat(t);var o=document.head||document.body;if(!o)throw new Error("Expected document.body not to be null. Stripe.js requires a <body> element.");return o.appendChild(n),n},Ne=function(e,t){!e||!e._registerWrapper||e._registerWrapper({name:"stripe-js",version:"2.4.0",startTime:t})},O=null,T=null,U=null,Le=function(e){return function(){e(new Error("Failed to load Stripe.js"))}},Ie=function(e,t){return function(){window.Stripe?e(window.Stripe):t(new Error("Stripe.js not available"))}},Te=function(e){return O!==null?O:(O=new Promise(function(t,n){if(typeof window>"u"||typeof document>"u"){t(null);return}if(window.Stripe&&e&&console.warn(X),window.Stripe){t(window.Stripe);return}try{var o=Ae();if(o&&e)console.warn(X);else if(!o)o=z(e);else if(o&&U!==null&&T!==null){var a;o.removeEventListener("load",U),o.removeEventListener("error",T),(a=o.parentNode)===null||a===void 0||a.removeChild(o),o=z(e)}U=Ie(t,n),T=Le(n),o.addEventListener("load",U),o.addEventListener("error",T)}catch(i){n(i);return}}),O.catch(function(t){return O=null,Promise.reject(t)}))},Ue=function(e,t,n){if(e===null)return null;var o=e.apply(void 0,t);return Ne(o,n),o},A,ue=!1,ie=function(){return A||(A=Te(null).catch(function(e){return A=null,Promise.reject(e)}),A)};Promise.resolve().then(function(){return ie()}).catch(function(r){ue||console.warn(r)});var We=function(){for(var e=arguments.length,t=new Array(e),n=0;n<e;n++)t[n]=arguments[n];ue=!0;var o=Date.now();return ie().then(function(a){return Ue(a,t,o)})};function H(r,e){var t=Object.keys(r);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(r);e&&(n=n.filter(function(o){return Object.getOwnPropertyDescriptor(r,o).enumerable})),t.push.apply(t,n)}return t}function Q(r){for(var e=1;e<arguments.length;e++){var t=arguments[e]!=null?arguments[e]:{};e%2?H(Object(t),!0).forEach(function(n){ce(r,n,t[n])}):Object.getOwnPropertyDescriptors?Object.defineProperties(r,Object.getOwnPropertyDescriptors(t)):H(Object(t)).forEach(function(n){Object.defineProperty(r,n,Object.getOwnPropertyDescriptor(t,n))})}return r}function W(r){"@babel/helpers - typeof";return typeof Symbol=="function"&&typeof Symbol.iterator=="symbol"?W=function(e){return typeof e}:W=function(e){return e&&typeof Symbol=="function"&&e.constructor===Symbol&&e!==Symbol.prototype?"symbol":typeof e},W(r)}function ce(r,e,t){return e in r?Object.defineProperty(r,e,{value:t,enumerable:!0,configurable:!0,writable:!0}):r[e]=t,r}function se(r,e){return Me(r)||_e(r,e)||Be(r,e)||qe()}function Me(r){if(Array.isArray(r))return r}function _e(r,e){var t=r&&(typeof Symbol<"u"&&r[Symbol.iterator]||r["@@iterator"]);if(t!=null){var n=[],o=!0,a=!1,i,l;try{for(t=t.call(r);!(o=(i=t.next()).done)&&(n.push(i.value),!(e&&n.length===e));o=!0);}catch(c){a=!0,l=c}finally{try{!o&&t.return!=null&&t.return()}finally{if(a)throw l}}return n}}function Be(r,e){if(r){if(typeof r=="string")return Z(r,e);var t=Object.prototype.toString.call(r).slice(8,-1);if(t==="Object"&&r.constructor&&(t=r.constructor.name),t==="Map"||t==="Set")return Array.from(r);if(t==="Arguments"||/^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(t))return Z(r,e)}}function Z(r,e){(e==null||e>r.length)&&(e=r.length);for(var t=0,n=new Array(e);t<e;t++)n[t]=r[t];return n}function qe(){throw new TypeError(`Invalid attempt to destructure non-iterable instance.
In order to be iterable, non-array objects must have a [Symbol.iterator]() method.`)}var $=function(e){var t=u.useRef(e);return u.useEffect(function(){t.current=e},[e]),t.current},w=function(e){return e!==null&&W(e)==="object"},$e=function(e){return w(e)&&typeof e.then=="function"},De=function(e){return w(e)&&typeof e.elements=="function"&&typeof e.createToken=="function"&&typeof e.createPaymentMethod=="function"&&typeof e.confirmCardPayment=="function"},ee="[object Object]",Ye=function r(e,t){if(!w(e)||!w(t))return e===t;var n=Array.isArray(e),o=Array.isArray(t);if(n!==o)return!1;var a=Object.prototype.toString.call(e)===ee,i=Object.prototype.toString.call(t)===ee;if(a!==i)return!1;if(!a&&!n)return e===t;var l=Object.keys(e),c=Object.keys(t);if(l.length!==c.length)return!1;for(var y={},h=0;h<l.length;h+=1)y[l[h]]=!0;for(var S=0;S<c.length;S+=1)y[c[S]]=!0;var C=Object.keys(y);if(C.length!==l.length)return!1;var g=e,k=t,E=function(P){return r(g[P],k[P])};return C.every(E)},Fe=function(e,t,n){return w(e)?Object.keys(e).reduce(function(o,a){var i=!w(t)||!Ye(e[a],t[a]);return n.includes(a)?(i&&console.warn("Unsupported prop change: options.".concat(a," is not a mutable property.")),o):i?Q(Q({},o||{}),{},ce({},a,e[a])):o},null):null},le="Invalid prop `stripe` supplied to `Elements`. We recommend using the `loadStripe` utility from `@stripe/stripe-js`. See https://stripe.com/docs/stripe-js/react#elements-props-stripe for details.",te=function(e){var t=arguments.length>1&&arguments[1]!==void 0?arguments[1]:le;if(e===null||De(e))return e;throw new Error(t)},Ke=function(e){var t=arguments.length>1&&arguments[1]!==void 0?arguments[1]:le;if($e(e))return{tag:"async",stripePromise:Promise.resolve(e).then(function(o){return te(o,t)})};var n=te(e,t);return n===null?{tag:"empty"}:{tag:"sync",stripe:n}},Ve=function(e){!e||!e._registerWrapper||!e.registerAppInfo||(e._registerWrapper({name:"react-stripe-js",version:"2.4.0"}),e.registerAppInfo({name:"react-stripe-js",version:"2.4.0",url:"https://stripe.com/docs/stripe-js/react"}))},de=u.createContext(null);de.displayName="CustomCheckoutSdkContext";var Ge=function(e,t){if(!e)throw new Error("Could not find CustomCheckoutProvider context; You need to wrap the part of your app that ".concat(t," in an <CustomCheckoutProvider> provider."));return e},Je=u.createContext(null);Je.displayName="CustomCheckoutContext";s.any,s.shape({clientSecret:s.string.isRequired,elementsOptions:s.object}).isRequired;var re=function(e){var t=u.useContext(de),n=u.useContext(fe);if(t&&n)throw new Error("You cannot wrap the part of your app that ".concat(e," in both <CustomCheckoutProvider> and <Elements> providers."));return t?Ge(t,e):Xe(n,e)},fe=u.createContext(null);fe.displayName="ElementsContext";var Xe=function(e,t){if(!e)throw new Error("Could not find Elements context; You need to wrap the part of your app that ".concat(t," in an <Elements> provider."));return e},pe=u.createContext(null);pe.displayName="CartElementContext";var ze=function(e,t){if(!e)throw new Error("Could not find Elements context; You need to wrap the part of your app that ".concat(t," in an <Elements> provider."));return e};s.any,s.object;var He={cart:null,cartState:null,setCart:function(){},setCartState:function(){}},ne=function(e){var t=arguments.length>1&&arguments[1]!==void 0?arguments[1]:!1,n=u.useContext(pe);return t?He:ze(n,e)};s.func.isRequired;var v=function(e,t,n){var o=!!n,a=u.useRef(n);u.useEffect(function(){a.current=n},[n]),u.useEffect(function(){if(!o||!e)return function(){};var i=function(){a.current&&a.current.apply(a,arguments)};return e.on(t,i),function(){e.off(t,i)}},[o,t,e,a])},Qe=function(e){return e.charAt(0).toUpperCase()+e.slice(1)},f=function(e,t){var n="".concat(Qe(e),"Element"),o=function(c){var y=c.id,h=c.className,S=c.options,C=S===void 0?{}:S,g=c.onBlur,k=c.onFocus,E=c.onReady,x=c.onChange,P=c.onEscape,he=c.onClick,ve=c.onLoadError,Ce=c.onLoaderStart,Ee=c.onNetworksChange,M=c.onCheckout,ye=c.onLineItemClick,Se=c.onConfirm,ge=c.onCancel,be=c.onShippingAddressChange,ke=c.onShippingRateChange,j=re("mounts <".concat(n,">")),N="elements"in j?j.elements:null,L="customCheckoutSdk"in j?j.customCheckoutSdk:null,xe=u.useState(null),Y=se(xe,2),m=Y[0],Pe=Y[1],b=u.useRef(null),_=u.useRef(null),F=ne("mounts <".concat(n,">"),"customCheckoutSdk"in j),B=F.setCart,q=F.setCartState;v(m,"blur",g),v(m,"focus",k),v(m,"escape",P),v(m,"click",he),v(m,"loaderror",ve),v(m,"loaderstart",Ce),v(m,"networkschange",Ee),v(m,"lineitemclick",ye),v(m,"confirm",Se),v(m,"cancel",ge),v(m,"shippingaddresschange",be),v(m,"shippingratechange",ke);var I;e==="cart"?I=function(V){q(V),E&&E(V)}:E&&(e==="expressCheckout"?I=E:I=function(){E(m)}),v(m,"ready",I);var we=e==="cart"?function(p){q(p),x&&x(p)}:x;v(m,"change",we);var je=e==="cart"?function(p){q(p),M&&M(p)}:M;v(m,"checkout",je),u.useLayoutEffect(function(){if(b.current===null&&_.current!==null&&(N||L)){var p=null;L?p=L.createElement(e,C):N&&(p=N.create(e,C)),e==="cart"&&B&&B(p),b.current=p,Pe(p),p&&p.mount(_.current)}},[N,L,C,B]);var K=$(C);return u.useEffect(function(){if(b.current){var p=Fe(C,K,["paymentRequest"]);p&&b.current.update(p)}},[C,K]),u.useLayoutEffect(function(){return function(){if(b.current&&typeof b.current.destroy=="function")try{b.current.destroy(),b.current=null}catch{}}},[]),u.createElement("div",{id:y,className:h,ref:_})},a=function(c){var y=re("mounts <".concat(n,">"));ne("mounts <".concat(n,">"),"customCheckoutSdk"in y);var h=c.id,S=c.className;return u.createElement("div",{id:h,className:S})},i=t?a:o;return i.propTypes={id:s.string,className:s.string,onChange:s.func,onBlur:s.func,onFocus:s.func,onReady:s.func,onEscape:s.func,onClick:s.func,onLoadError:s.func,onLoaderStart:s.func,onNetworksChange:s.func,onCheckout:s.func,onLineItemClick:s.func,onConfirm:s.func,onCancel:s.func,onShippingAddressChange:s.func,onShippingRateChange:s.func,options:s.object},i.displayName=n,i.__elementType=e,i},d=typeof window>"u",D=u.createContext(null);D.displayName="EmbeddedCheckoutProviderContext";var me=function(){var e=u.useContext(D);if(!e)throw new Error("<EmbeddedCheckout> must be used within <EmbeddedCheckoutProvider>");return e},Ze="Invalid prop `stripe` supplied to `EmbeddedCheckoutProvider`. We recommend using the `loadStripe` utility from `@stripe/stripe-js`. See https://stripe.com/docs/stripe-js/react#elements-props-stripe for details.",et=function(e){var t=e.stripe,n=e.options,o=e.children,a=u.useMemo(function(){return Ke(t,Ze)},[t]),i=u.useRef(null),l=u.useRef(null),c=u.useState({embeddedCheckout:null}),y=se(c,2),h=y[0],S=y[1];u.useEffect(function(){if(!(l.current||i.current)){var k=function(x){l.current||i.current||(l.current=x,i.current=l.current.initEmbeddedCheckout(n).then(function(P){S({embeddedCheckout:P})}))};a.tag==="async"&&!l.current&&n.clientSecret?a.stripePromise.then(function(E){E&&k(E)}):a.tag==="sync"&&!l.current&&n.clientSecret&&k(a.stripe)}},[a,n,h,l]),u.useEffect(function(){return function(){h.embeddedCheckout?(i.current=null,h.embeddedCheckout.destroy()):i.current&&i.current.then(function(){i.current=null,h.embeddedCheckout&&h.embeddedCheckout.destroy()})}},[h.embeddedCheckout]),u.useEffect(function(){Ve(l)},[l]);var C=$(t);u.useEffect(function(){C!==null&&C!==t&&console.warn("Unsupported prop change on EmbeddedCheckoutProvider: You cannot change the `stripe` prop after setting it.")},[C,t]);var g=$(n);return u.useEffect(function(){if(g!=null){if(n==null){console.warn("Unsupported prop change on EmbeddedCheckoutProvider: You cannot unset options after setting them.");return}g.clientSecret!=null&&n.clientSecret!==g.clientSecret&&console.warn("Unsupported prop change on EmbeddedCheckoutProvider: You cannot change the client secret after setting it. Unmount and create a new instance of EmbeddedCheckoutProvider instead."),g.onComplete!=null&&n.onComplete!==g.onComplete&&console.warn("Unsupported prop change on EmbeddedCheckoutProvider: You cannot change the onComplete option after setting it.")}},[g,n]),u.createElement(D.Provider,{value:h},o)},tt=function(e){var t=e.id,n=e.className,o=me(),a=o.embeddedCheckout,i=u.useRef(!1),l=u.useRef(null);return u.useLayoutEffect(function(){return!i.current&&a&&l.current!==null&&(a.mount(l.current),i.current=!0),function(){if(i.current&&a)try{a.unmount(),i.current=!1}catch{}}},[a]),u.createElement("div",{ref:l,id:t,className:n})},rt=function(e){var t=e.id,n=e.className;return me(),u.createElement("div",{id:t,className:n})},nt=d?rt:tt;f("auBankAccount",d);f("card",d);f("cardNumber",d);f("cardExpiry",d);f("cardCvc",d);f("fpxBank",d);f("iban",d);f("idealBank",d);f("p24Bank",d);f("epsBank",d);f("payment",d);f("expressCheckout",d);f("paymentRequestButton",d);f("linkAuthentication",d);f("address",d);f("shippingAddress",d);f("cart",d);f("paymentMethodMessaging",d);f("affirmMessage",d);f("afterpayClearpayMessage",d);const ot=We(oe.stripePublicKey),ut=()=>{const[r,e]=G.useState(""),t=Re().id,n=new URL(`/api/checkout/create-checkout-session/${t}`,oe.baseUrl);return G.useEffect(()=>{fetch(n,{method:"POST"}).then(o=>o.json()).then(o=>{e(o.session.client_secret)}).catch(o=>{console.error("Error fetching checkout session:",o)})},[]),R.jsxs("div",{className:J.checkout,children:[R.jsx("h1",{className:J.subHeading,children:"Checkout Form:"}),R.jsx("div",{id:"checkout",children:r&&R.jsx(et,{stripe:ot,options:{clientSecret:r},children:R.jsx(nt,{})})})]})};export{ut as default};
