import { Component, For } from 'solid-js';
import styles from '../styles/Suppliers.module.css';
import utils from '../styles/Utils.module.css';
import { Video } from '../types/video';

interface Props {
    videos: Video[];
}

const Videos: Component<Props> = props => {
    return (
        <section id="videos">
            <h1 class={utils.sectionHeading}>Videos</h1>
            <div class={styles.videoGrid}>
                <For each={props.videos}>
                    {video => (
                        <div class={styles.videoCard} id={video.title || ''}>
                            <h1 class={utils.mainHeading}>{video.title}</h1>
                            <iframe
                                width="425"
                                height="315"
                                class={styles.video}
                                src={`https://www.youtube.com/embed/${video.video_id}`}
                                allowfullscreen
                            />
                        </div>
                    )}
                </For>
            </div>
        </section>
    );
};

export default Videos;
