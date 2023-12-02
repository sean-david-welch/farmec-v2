import { onMount, createSignal } from 'solid-js';
import Typewriter from 'typewriter-effect';
import utils from '../../styles/Utils.module.css';

const TypewriterComponent = () => {
  const [isClient, setIsClient] = createSignal(false);
  let typewriterElement: HTMLHeadingElement | undefined;

  onMount(() => {
    setIsClient(true);

    if (typewriterElement) {
      const typewriter = new Typewriter(typewriterElement, {
        loop: false,
        cursor: '',
        delay: 50,
      });

      typewriter.stop().typeString('Importers & Distributors of Quality Agricultural Machinery').start();
    }
  });

  return (
    <>
      {!isClient() ? (
        <div class={utils.typewriterSkeleton}></div>
      ) : (
        <div class={utils.typewriter}>
          <h1
            ref={el => {
              typewriterElement = el;
            }}></h1>
          <button class={utils.btn}>
            <a href="#Info">
              Find Out More:
              {/* FontAwesomeIcon equivalent */}
            </a>
          </button>
        </div>
      )}
    </>
  );
};

export default TypewriterComponent;
