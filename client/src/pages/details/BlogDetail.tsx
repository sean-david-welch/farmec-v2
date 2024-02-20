import styles from '../../styles/Blogs.module.css';
import utils from '../../styles/Utils.module.css';

import ErrorPage from '../../layouts/Error';
import Loading from '../../layouts/Loading';
import BlogForm from '../../forms/BlogForm';
import DeleteButton from '../../components/DeleteButton';

import { Blog } from '../../types/blogTypes';
import { useParams } from 'react-router-dom';
import { useGetResourceById } from '../../hooks/genericHooks';
import { Fragment, useEffect } from 'react';
import { useUserStore } from '../../lib/store';

const BlogDetail: React.FC = () => {
    const id = useParams<{ id: string }>().id as string;

    const { isAdmin } = useUserStore();
    const { data: blog, isLoading, isError } = useGetResourceById<Blog>('blogs', id);

    useEffect(() => {}, [id]);

    if (isError) return <ErrorPage />;
    if (isLoading) return <Loading />;

    return (
        <section id="blog">
            <Fragment>
                {blog && (
                    <div className={styles.blogDetail}>
                        <h1 className={utils.sectionHeading}>{blog.title}</h1>
                        <div className={styles.blogBody}>
                            <img src={blog.main_image || '/default.jpg'} alt="Blog image" width={600} height={600} />
                            <h1 className={utils.mainHeading}>{blog.subheading}</h1>
                            <p className={utils.paragraph}>{blog.body}</p>
                        </div>
                    </div>
                )}
                {isAdmin && blog?.id && (
                    <div className={utils.optionsBtn}>
                        <BlogForm id={blog?.id} blog={blog} />
                        <DeleteButton id={blog?.id} resourceKey="blogs" navigateBack={true} />
                    </div>
                )}
            </Fragment>
        </section>
    );
};

export default BlogDetail;
