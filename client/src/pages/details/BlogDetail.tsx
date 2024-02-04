import styles from '../../styles/Blogs.module.css';
import utils from '../../styles/Utils.module.css';

import { Blog } from '../../types/blogTypes';
import { useParams } from 'react-router-dom';
import { useGetResourceById } from '../../hooks/genericHooks';

const BlogDetail: React.FC = () => {
    const id = useParams<{ id: string }>().id as string;
    const { data: blog, isLoading } = useGetResourceById<Blog>('blogs', id);

    if (isLoading) {
        return <div>Loading...</div>;
    }

    return (
        <section id="blog">
            {blog && (
                <div className={styles.blogDetail}>
                    <h1 className={utils.sectionHeading}>{blog.title}</h1>
                    <div className={styles.blogBody}>
                        <img
                            src={blog.main_image || '/default.jpg'}
                            alt={'/default.jpg'}
                            width={400}
                            height={400}
                        />
                        <h1 className={utils.mainHeading}>{blog.subheading}</h1>
                        <p className={utils.paragraph}>{blog.body}</p>
                    </div>
                </div>
            )}
        </section>
    );
};

export default BlogDetail;
