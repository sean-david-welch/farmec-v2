import styles from '../../styles/Blogs.module.css';
import utils from '../../styles/Utils.module.css';

import { Blog } from '../../types/blogTypes';
import { useParams } from 'react-router-dom';
import { useGetResourceById } from '../../hooks/genericHooks';
import { Fragment } from 'react';
import { useUserStore } from '../../lib/store';
import BlogForm from '../../forms/BlogForm';
import DeleteButton from '../../components/DeleteButton';

const BlogDetail: React.FC = () => {
    const { isAdmin } = useUserStore();
    const id = useParams<{ id: string }>().id as string;
    const { data: blog, isLoading } = useGetResourceById<Blog>('blogs', id);

    if (isLoading) {
        return <div>Loading...</div>;
    }

    return (
        <section id="blog">
            <Fragment>
                {blog && (
                    <div className={styles.blogDetail}>
                        <h1 className={utils.sectionHeading}>{blog.title}</h1>
                        <div className={styles.blogBody}>
                            <img
                                src={blog.main_image || '/default.jpg'}
                                alt="Blog image"
                                width={600}
                                height={600}
                            />
                            <h1 className={utils.mainHeading}>{blog.subheading}</h1>
                            <p className={utils.paragraph}>{blog.body}</p>
                        </div>
                    </div>
                )}
                {isAdmin && blog.id && (
                    <div className={utils.optionsBtn}>
                        <BlogForm id={blog.id} blog={blog} />
                        <DeleteButton id={blog.id} resourceKey="blogs" />
                    </div>
                )}
            </Fragment>
        </section>
    );
};

export default BlogDetail;
