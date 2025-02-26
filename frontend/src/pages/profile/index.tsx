import { Suspense } from 'solid-js';
import ProfileData from '@/pages/profile/profile.data';
import Title from '@/components/title';

export default function ProfilePage() {
    const name = ProfileData();

    return (
        <>
            <Title>Profile</Title>
            <section class="bg-pink-100 text-gray-700 p-8">
                <h1 class="text-2xl font-bold">Profile</h1>

                <p class="mt-4">Profile page</p>

                <p>
                    <span>We love</span>
                    <Suspense fallback={<span>...</span>}>
                        <span>&nbsp;{name()}</span>
                    </Suspense>
                </p>
            </section>
        </>
    );
}