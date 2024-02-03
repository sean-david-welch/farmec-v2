import DeleteButton from '../components/DeleteButton';
import VideoForm from '../forms/VideoForm';
import styles from '../styles/Suppliers.module.css';
import utils from '../styles/Utils.module.css';
import { Supplier } from '../types/supplierTypes';

import { Video } from '../types/videoTypes';

interface VideoProps {
    videos: Video[];
    isAdmin: boolean;
    suppliers: Supplier[];
}

const Videos: React.FC<VideoProps> = ({ videos, isAdmin, suppliers }) => {
    return (
        <section id="videos">
            <h1 className={utils.sectionHeading}>Videos</h1>
            {isAdmin && <VideoForm suppliers={suppliers} />}

            <div className={styles.videoGrid}>
                {videos.map((video) => (
                    <div className={styles.videoCard} id={video.title || ''} key={video.id}>
                        <h1 className={utils.mainHeading}>{video.title}</h1>
                        <iframe
                            width="425"
                            height="315"
                            className={styles.video}
                            src={`https://www.youtube.com/embed/${video.video_id}`}
                            allowFullScreen
                        />
                        {isAdmin && video.id && (
                            <div className={utils.optionsBtn}>
                                <VideoForm id={video.id} suppliers={suppliers} video={video} />
                                <DeleteButton id={video.id} resourceKey="videos" />
                            </div>
                        )}
                    </div>
                ))}
            </div>
        </section>
    );
};

export default Videos;
