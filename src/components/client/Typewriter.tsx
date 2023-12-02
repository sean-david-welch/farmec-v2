import { onMount } from 'solid-js';
import Typewriter from 'typewriter-effect';
import utils from '../../styles/Utils.module.css';

const TypewriterComponent = () => {
  let typewriterElement: HTMLElement | null = null;

  onMount(() => {
    if (typewriterElement) {
      const typewriter = new Typewriter(typewriterElement, {
        loop: true,
        delay: 75,
      });

      if (typeof typewriter.typeString === 'function') {
        typewriter.typeString('Importers & Distributors of Quality Agricultural Machinery').start();
      } else {
        console.error('typeString method is not available on typewriter instance');
      }
    }
  });

  return (
    <div class={utils.typewriter}>
      <h1>Importers & Distributors of Quality Agricultural Machinery</h1>

      <button class={utils.btn}>
        <a href="#Info">
          Find Out More: <img src="/icons/chevron-down.svg" alt="down" />
        </a>
      </button>
    </div>
  );
};

export default TypewriterComponent;
