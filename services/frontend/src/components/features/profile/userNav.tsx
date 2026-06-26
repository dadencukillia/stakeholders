import { Button } from "@/components/ui/button";
import {
	DropdownMenu,
	DropdownMenuContent,
	DropdownMenuGroup,
	DropdownMenuItem,
	DropdownMenuLabel,
	DropdownMenuSeparator,
	DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { Skeleton } from "@/components/ui/skeleton";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import useSWR from "swr";
import { fetcher } from "@/lib/fetcher";

export const UserNavData = ({
	avatarUrl,
	username,
}: {
	avatarUrl: string;
	username: string;
}) => {
	return (
		<DropdownMenu>
			<DropdownMenuTrigger asChild>
				<Button variant="ghost" className="px-4 py-6">
					<div className="flex flex-row gap-2 items-center">
						<Avatar>
							<AvatarImage src={avatarUrl} alt="user avatar" />
							<AvatarFallback>A</AvatarFallback>
						</Avatar>
						<div className="max-w-48 overflow-hidden text-ellipsis hidden sm:block text-nowrap">
							{username}
						</div>
					</div>
				</Button>
			</DropdownMenuTrigger>
			<DropdownMenuContent>
				<DropdownMenuGroup>
					<DropdownMenuLabel>My Account</DropdownMenuLabel>
					<DropdownMenuItem>Profile</DropdownMenuItem>
					<DropdownMenuItem>Billing</DropdownMenuItem>
				</DropdownMenuGroup>
				<DropdownMenuSeparator />
				<DropdownMenuGroup>
					<DropdownMenuItem>Team</DropdownMenuItem>
					<DropdownMenuItem>Subscription</DropdownMenuItem>
				</DropdownMenuGroup>
			</DropdownMenuContent>
		</DropdownMenu>
	);
};

export const UserNavUnlogged = () => {
	return (
		<div className="flex items-center gap-2">
			<Button variant="ghost" size="sm">
				Log in
			</Button>
			<Button size="sm">Get started</Button>
		</div>
	);
};

export const UserNavPlaceholder = () => {
	return (
		<div className="flex flex-row gap-2 items-center px-4 py-6">
			<Skeleton className="rounded-full aspect-square size-8" />
			<Skeleton className="w-24 h-3 hidden sm:block" />
		</div>
	);
};

export const UserNavDynamic = () => {
	const { data, error, isLoading } = useSWR(
		"https://www.fakerapi.it/api/v2/users?_quantity=1",
		fetcher,
		{
			revalidateOnFocus: false,
			revalidateOnReconnect: false,
		},
	);

	if (error)
		return (
			<div className="flex flex-row gap-2 items-center px-4 py-6">
				<div className="rounded-full aspect-square size-8 bg-destructive/80" />
				<div className="rounded-full w-24 h-3 bg-destructive/80 hidden sm:block" />
			</div>
		);
	if (isLoading) return <UserNavPlaceholder />;

	return (
		<UserNavData
			avatarUrl={data.data[0].image}
			username={data.data[0].username}
		/>
	);
};
