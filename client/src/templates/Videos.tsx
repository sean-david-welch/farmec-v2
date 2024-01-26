import styles from '../styles/Suppliers.module.css';
import utils from '../styles/Utils.module.css';

import { Video } from '../types/video';

interface VideoProps {
    videos: Video[];
}

const Videos: React.FC<VideoProps> = props => {
    return (
        <section id="videos">
            <h1 className={utils.sectionHeading}>Videos</h1>
            <div className={styles.videoGrid}>
                {props.videos.map(video => (
                    <div className={styles.videoCard} id={video.title || ''}>
                        <h1 className={utils.mainHeading}>{video.title}</h1>
                        <iframe
                            width="425"
                            height="315"
                            className={styles.video}
                            src={`https://www.youtube.com/embed/${video.video_id}`}
                            allowFullScreen
                        />
                    </div>
                ))}
            </div>
        </section>
    );
};

export default Videos;
