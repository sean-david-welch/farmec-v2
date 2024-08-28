import styles from '../../styles/Blogs.module.css';
import utils from '../../styles/Utils.module.css';

import ErrorPage from '../../layouts/Error';
import Loading from '../../layouts/Loading';
import BlogForm from '../../forms/BlogForm';
import DeleteButton from '../../components/DeleteButton';

import { Blog } from '../../types/blogTypes';
import { useParams } from 'react-router-dom';
import { useGetResourceById } from '../../hooks/genericHooks';
import {FC, Fragment, SyntheticEvent, useEffect} from 'react';
import { useUserStore } from '../../lib/store';
import { Helmet } from 'react-helmet';

const BlogDetail: FC = () => {
	const id = useParams<{ id: string }>().id as string;

	const { isAdmin } = useUserStore();
	const { data: blog, isLoading, isError } = useGetResourceById<Blog>('blogs', id);

	useEffect(() => {}, [id]);

	if (isError) return <ErrorPage />;
	if (isLoading) return <Loading />;

	const imageError = (event: SyntheticEvent<HTMLImageElement, Event>) => {
		event.currentTarget.src = '/default.jpg';
	};

	return (
		<>
			{blog && (
				<Helmet>
					<title>{`${blog.title} - Farmec Blog`}</title>
					<meta name="description" content={blog.subheading} />

					<meta property="og:title" content={`${blog.title} - Farmec Blog`} />
					<meta property="og:description" content={blog.subheading} />
					<meta property="og:image" content="https://www.farmec.ie/farmec_images/Suppliers/sip1250.webp" />
					<meta property="og:url" content={`https://www.farmec.ie/blogs/${blog.id}`} />
					<meta property="og:type" content="article" />

					<meta name="twitter:card" content="summary_large_image" />
					<meta name="twitter:title" content={`${blog.title} - Farmec Blog`} />
					<meta name="twitter:description" content={blog.subheading} />
					<meta name="twitter:image" content="https://www.farmec.ie/farmec_images/Suppliers/sip1250.webp" />

					<link rel="canonical" href={`https://www.farmec.ie/blogs/${blog.id}`} />
				</Helmet>
			)}

			<section id="blog">
				<Fragment>
					{blog && (
						<div className={styles.blogDetail}>
							<h1 className={utils.sectionHeading}>{blog.title}</h1>
							<div className={styles.blogBody}>
								<img src={blog.main_image} alt="Blog image" width={600} height={600} onError={imageError} />
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
		</>
	);
};

export default BlogDetail;
