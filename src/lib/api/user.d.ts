// Strip prisma user object type of password, createdAt, updatedAt,

type CurrentUserResponse = {
    id: string;
    name: string;
    email: string;
    bio: string?;
};
