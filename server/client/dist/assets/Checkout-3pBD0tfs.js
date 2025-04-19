import{R as u,P as s,r as G,y as je,c as ne,j as R,u as J}from"./index-yDAfzZuS.js";var oe="https://js.stripe.com/v3",Re=/^https:\/\/js\.stripe\.com\/v3\/?(\?.*)?$/;var Oe=function(){for(var e=document.querySelectorAll('script[src^="'.concat(oe,'"]')),t=0;t<e.length;t++){var n=e[t];if(Re.test(n.src))return n}return null},X=function(e){var t="",n=document.createElement("script");n.src="".concat(oe).concat(t);var o=document.head||document.body;if(!o)throw new Error("Expected document.body not to be null. Stripe.js requires a <body> element.");return o.appendChild(n),n},Ae=function(e,t){!e||!e._registerWrapper||e._registerWrapper({name:"stripe-js",version:"2.4.0",startTime:t})},O=null,T=null,U=null,Ne=function(e){return function(){e(new Error("Failed to load Stripe.js"))}},Le=function(e,t){return function(){window.Stripe?e(window.Stripe):t(new Error("Stripe.js not available"))}},Ie=function(e){return O!==null?O:(O=new Promise(function(t,n){if(typeof window>"u"||typeof document>"u"){t(null);return}if(window.Stripe){t(window.Stripe);return}try{var o=Oe();if(!(o&&e)){if(!o)o=X(e);else if(o&&U!==null&&T!==null){var a;o.removeEventListener("load",U),o.removeEventListener("error",T),(a=o.parentNode)===null||a===void 0||a.removeChild(o),o=X(e)}}U=Le(t,n),T=Ne(n),o.addEventListener("load",U),o.addEventListener("error",T)}catch(i){n(i);return}}),O.catch(function(t){return O=null,Promise.reject(t)}))},Te=function(e,t,n){if(e===null)return null;var o=e.apply(void 0,t);return Ae(o,n),o},A,ae=!1,ue=function(){return A||(A=Ie(null).catch(function(e){return A=null,Promise.reject(e)}),A)};Promise.resolve().then(function(){return ue()}).catch(function(r){ae||console.warn(r)});var Ue=function(){for(var e=arguments.length,t=new Array(e),n=0;n<e;n++)t[n]=arguments[n];ae=!0;var o=Date.now();return ue().then(function(a){return Te(a,t,o)})};function z(r,e){var t=Object.keys(r);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(r);e&&(n=n.filter(function(o){return Object.getOwnPropertyDescriptor(r,o).enumerable})),t.push.apply(t,n)}return t}function H(r){for(var e=1;e<arguments.length;e++){var t=arguments[e]!=null?arguments[e]:{};e%2?z(Object(t),!0).forEach(function(n){ie(r,n,t[n])}):Object.getOwnPropertyDescriptors?Object.defineProperties(r,Object.getOwnPropertyDescriptors(t)):z(Object(t)).forEach(function(n){Object.defineProperty(r,n,Object.getOwnPropertyDescriptor(t,n))})}return r}function W(r){"@babel/helpers - typeof";return typeof Symbol=="function"&&typeof Symbol.iterator=="symbol"?W=function(e){return typeof e}:W=function(e){return e&&typeof Symbol=="function"&&e.constructor===Symbol&&e!==Symbol.prototype?"symbol":typeof e},W(r)}function ie(r,e,t){return e in r?Object.defineProperty(r,e,{value:t,enumerable:!0,configurable:!0,writable:!0}):r[e]=t,r}function ce(r,e){return We(r)||Me(r,e)||_e(r,e)||Be()}function We(r){if(Array.isArray(r))return r}function Me(r,e){var t=r&&(typeof Symbol<"u"&&r[Symbol.iterator]||r["@@iterator"]);if(t!=null){var n=[],o=!0,a=!1,i,l;try{for(t=t.call(r);!(o=(i=t.next()).done)&&(n.push(i.value),!(e&&n.length===e));o=!0);}catch(c){a=!0,l=c}finally{try{!o&&t.return!=null&&t.return()}finally{if(a)throw l}}return n}}function _e(r,e){if(r){if(typeof r=="string")return Q(r,e);var t=Object.prototype.toString.call(r).slice(8,-1);if(t==="Object"&&r.constructor&&(t=r.constructor.name),t==="Map"||t==="Set")return Array.from(r);if(t==="Arguments"||/^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(t))return Q(r,e)}}function Q(r,e){(e==null||e>r.length)&&(e=r.length);for(var t=0,n=new Array(e);t<e;t++)n[t]=r[t];return n}function Be(){throw new TypeError(`Invalid attempt to destructure non-iterable instance.
In order to be iterable, non-array objects must have a [Symbol.iterator]() method.`)}var $=function(e){var t=u.useRef(e);return u.useEffect(function(){t.current=e},[e]),t.current},w=function(e){return e!==null&&W(e)==="object"},qe=function(e){return w(e)&&typeof e.then=="function"},$e=function(e){return w(e)&&typeof e.elements=="function"&&typeof e.createToken=="function"&&typeof e.createPaymentMethod=="function"&&typeof e.confirmCardPayment=="function"},Z="[object Object]",De=function r(e,t){if(!w(e)||!w(t))return e===t;var n=Array.isArray(e),o=Array.isArray(t);if(n!==o)return!1;var a=Object.prototype.toString.call(e)===Z,i=Object.prototype.toString.call(t)===Z;if(a!==i)return!1;if(!a&&!n)return e===t;var l=Object.keys(e),c=Object.keys(t);if(l.length!==c.length)return!1;for(var y={},h=0;h<l.length;h+=1)y[l[h]]=!0;for(var g=0;g<c.length;g+=1)y[c[g]]=!0;var C=Object.keys(y);if(C.length!==l.length)return!1;var S=e,k=t,E=function(P){return r(S[P],k[P])};return C.every(E)},Ye=function(e,t,n){return w(e)?Object.keys(e).reduce(function(o,a){var i=!w(t)||!De(e[a],t[a]);return n.includes(a)?(i&&console.warn("Unsupported prop change: options.".concat(a," is not a mutable property.")),o):i?H(H({},o||{}),{},ie({},a,e[a])):o},null):null},se="Invalid prop `stripe` supplied to `Elements`. We recommend using the `loadStripe` utility from `@stripe/stripe-js`. See https://stripe.com/docs/stripe-js/react#elements-props-stripe for details.",ee=function(e){var t=arguments.length>1&&arguments[1]!==void 0?arguments[1]:se;if(e===null||$e(e))return e;throw new Error(t)},Fe=function(e){var t=arguments.length>1&&arguments[1]!==void 0?arguments[1]:se;if(qe(e))return{tag:"async",stripePromise:Promise.resolve(e).then(function(o){return ee(o,t)})};var n=ee(e,t);return n===null?{tag:"empty"}:{tag:"sync",stripe:n}},Ke=function(e){!e||!e._registerWrapper||!e.registerAppInfo||(e._registerWrapper({name:"react-stripe-js",version:"2.4.0"}),e.registerAppInfo({name:"react-stripe-js",version:"2.4.0",url:"https://stripe.com/docs/stripe-js/react"}))},le=u.createContext(null);le.displayName="CustomCheckoutSdkContext";var Ve=function(e,t){if(!e)throw new Error("Could not find CustomCheckoutProvider context; You need to wrap the part of your app that ".concat(t," in an <CustomCheckoutProvider> provider."));return e},Ge=u.createContext(null);Ge.displayName="CustomCheckoutContext";s.any,s.shape({clientSecret:s.string.isRequired,elementsOptions:s.object}).isRequired;var te=function(e){var t=u.useContext(le),n=u.useContext(de);if(t&&n)throw new Error("You cannot wrap the part of your app that ".concat(e," in both <CustomCheckoutProvider> and <Elements> providers."));return t?Ve(t,e):Je(n,e)},de=u.createContext(null);de.displayName="ElementsContext";var Je=function(e,t){if(!e)throw new Error("Could not find Elements context; You need to wrap the part of your app that ".concat(t," in an <Elements> provider."));return e},fe=u.createContext(null);fe.displayName="CartElementContext";var Xe=function(e,t){if(!e)throw new Error("Could not find Elements context; You need to wrap the part of your app that ".concat(t," in an <Elements> provider."));return e};s.any,s.object;var ze={cart:null,cartState:null,setCart:function(){},setCartState:function(){}},re=function(e){var t=arguments.length>1&&arguments[1]!==void 0?arguments[1]:!1,n=u.useContext(fe);return t?ze:Xe(n,e)};s.func.isRequired;var v=function(e,t,n){var o=!!n,a=u.useRef(n);u.useEffect(function(){a.current=n},[n]),u.useEffect(function(){if(!o||!e)return function(){};var i=function(){a.current&&a.current.apply(a,arguments)};return e.on(t,i),function(){e.off(t,i)}},[o,t,e,a])},He=function(e){return e.charAt(0).toUpperCase()+e.slice(1)},f=function(e,t){var n="".concat(He(e),"Element"),o=function(c){var y=c.id,h=c.className,g=c.options,C=g===void 0?{}:g,S=c.onBlur,k=c.onFocus,E=c.onReady,x=c.onChange,P=c.onEscape,me=c.onClick,he=c.onLoadError,ve=c.onLoaderStart,Ce=c.onNetworksChange,M=c.onCheckout,Ee=c.onLineItemClick,ye=c.onConfirm,ge=c.onCancel,Se=c.onShippingAddressChange,be=c.onShippingRateChange,j=te("mounts <".concat(n,">")),N="elements"in j?j.elements:null,L="customCheckoutSdk"in j?j.customCheckoutSdk:null,ke=u.useState(null),Y=ce(ke,2),m=Y[0],xe=Y[1],b=u.useRef(null),_=u.useRef(null),F=re("mounts <".concat(n,">"),"customCheckoutSdk"in j),B=F.setCart,q=F.setCartState;v(m,"blur",S),v(m,"focus",k),v(m,"escape",P),v(m,"click",me),v(m,"loaderror",he),v(m,"loaderstart",ve),v(m,"networkschange",Ce),v(m,"lineitemclick",Ee),v(m,"confirm",ye),v(m,"cancel",ge),v(m,"shippingaddresschange",Se),v(m,"shippingratechange",be);var I;e==="cart"?I=function(V){q(V),E&&E(V)}:E&&(e==="expressCheckout"?I=E:I=function(){E(m)}),v(m,"ready",I);var Pe=e==="cart"?function(p){q(p),x&&x(p)}:x;v(m,"change",Pe);var we=e==="cart"?function(p){q(p),M&&M(p)}:M;v(m,"checkout",we),u.useLayoutEffect(function(){if(b.current===null&&_.current!==null&&(N||L)){var p=null;L?p=L.createElement(e,C):N&&(p=N.create(e,C)),e==="cart"&&B&&B(p),b.current=p,xe(p),p&&p.mount(_.current)}},[N,L,C,B]);var K=$(C);return u.useEffect(function(){if(b.current){var p=Ye(C,K,["paymentRequest"]);p&&b.current.update(p)}},[C,K]),u.useLayoutEffect(function(){return function(){if(b.current&&typeof b.current.destroy=="function")try{b.current.destroy(),b.current=null}catch{}}},[]),u.createElement("div",{id:y,className:h,ref:_})},a=function(c){var y=te("mounts <".concat(n,">"));re("mounts <".concat(n,">"),"customCheckoutSdk"in y);var h=c.id,g=c.className;return u.createElement("div",{id:h,className:g})},i=t?a:o;return i.propTypes={id:s.string,className:s.string,onChange:s.func,onBlur:s.func,onFocus:s.func,onReady:s.func,onEscape:s.func,onClick:s.func,onLoadError:s.func,onLoaderStart:s.func,onNetworksChange:s.func,onCheckout:s.func,onLineItemClick:s.func,onConfirm:s.func,onCancel:s.func,onShippingAddressChange:s.func,onShippingRateChange:s.func,options:s.object},i.displayName=n,i.__elementType=e,i},d=typeof window>"u",D=u.createContext(null);D.displayName="EmbeddedCheckoutProviderContext";var pe=function(){var e=u.useContext(D);if(!e)throw new Error("<EmbeddedCheckout> must be used within <EmbeddedCheckoutProvider>");return e},Qe="Invalid prop `stripe` supplied to `EmbeddedCheckoutProvider`. We recommend using the `loadStripe` utility from `@stripe/stripe-js`. See https://stripe.com/docs/stripe-js/react#elements-props-stripe for details.",Ze=function(e){var t=e.stripe,n=e.options,o=e.children,a=u.useMemo(function(){return Fe(t,Qe)},[t]),i=u.useRef(null),l=u.useRef(null),c=u.useState({embeddedCheckout:null}),y=ce(c,2),h=y[0],g=y[1];u.useEffect(function(){if(!(l.current||i.current)){var k=function(x){l.current||i.current||(l.current=x,i.current=l.current.initEmbeddedCheckout(n).then(function(P){g({embeddedCheckout:P})}))};a.tag==="async"&&!l.current&&n.clientSecret?a.stripePromise.then(function(E){E&&k(E)}):a.tag==="sync"&&!l.current&&n.clientSecret&&k(a.stripe)}},[a,n,h,l]),u.useEffect(function(){return function(){h.embeddedCheckout?(i.current=null,h.embeddedCheckout.destroy()):i.current&&i.current.then(function(){i.current=null,h.embeddedCheckout&&h.embeddedCheckout.destroy()})}},[h.embeddedCheckout]),u.useEffect(function(){Ke(l)},[l]);var C=$(t);u.useEffect(function(){C!==null&&C!==t&&console.warn("Unsupported prop change on EmbeddedCheckoutProvider: You cannot change the `stripe` prop after setting it.")},[C,t]);var S=$(n);return u.useEffect(function(){if(S!=null){if(n==null){console.warn("Unsupported prop change on EmbeddedCheckoutProvider: You cannot unset options after setting them.");return}S.clientSecret!=null&&n.clientSecret!==S.clientSecret&&console.warn("Unsupported prop change on EmbeddedCheckoutProvider: You cannot change the client secret after setting it. Unmount and create a new instance of EmbeddedCheckoutProvider instead."),S.onComplete!=null&&n.onComplete!==S.onComplete&&console.warn("Unsupported prop change on EmbeddedCheckoutProvider: You cannot change the onComplete option after setting it.")}},[S,n]),u.createElement(D.Provider,{value:h},o)},et=function(e){var t=e.id,n=e.className,o=pe(),a=o.embeddedCheckout,i=u.useRef(!1),l=u.useRef(null);return u.useLayoutEffect(function(){return!i.current&&a&&l.current!==null&&(a.mount(l.current),i.current=!0),function(){if(i.current&&a)try{a.unmount(),i.current=!1}catch{}}},[a]),u.createElement("div",{ref:l,id:t,className:n})},tt=function(e){var t=e.id,n=e.className;return pe(),u.createElement("div",{id:t,className:n})},rt=d?tt:et;f("auBankAccount",d);f("card",d);f("cardNumber",d);f("cardExpiry",d);f("cardCvc",d);f("fpxBank",d);f("iban",d);f("idealBank",d);f("p24Bank",d);f("epsBank",d);f("payment",d);f("expressCheckout",d);f("paymentRequestButton",d);f("linkAuthentication",d);f("address",d);f("shippingAddress",d);f("cart",d);f("paymentMethodMessaging",d);f("affirmMessage",d);f("afterpayClearpayMessage",d);const nt=Ue(ne.stripePublicKey),at=()=>{const[r,e]=G.useState(""),t=je().id,n=new URL(`/api/checkout/create-checkout-session/${t}`,ne.baseUrl);return G.useEffect(()=>{fetch(n,{method:"POST"}).then(o=>o.json()).then(o=>{e(o.session.client_secret)}).catch(o=>{console.error("Error fetching checkout session:",o)})},[]),R.jsxs("div",{className:J.checkout,children:[R.jsx("h1",{className:J.subHeading,children:"Checkout Form:"}),R.jsx("div",{id:"checkout",children:r&&R.jsx(Ze,{stripe:nt,options:{clientSecret:r},children:R.jsx(rt,{})})})]})};export{at as default};
