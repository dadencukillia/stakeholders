const localStorageKey = "darkTheme";

export const getUserPreferedTheme = () => {
	const localTheme = localStorage.getItem(localStorageKey);
	if (!localTheme) {
		return undefined;
	}

	return localTheme === "true" ? "dark" : "light";
};

export const getBrowserTheme = () => {
	return window.matchMedia("(prefers-color-scheme: dark)").matches
		? "dark"
		: "light";
};

export const getTheme = () => {
	const userPrefered = getUserPreferedTheme();

	if (!userPrefered) return getBrowserTheme();

	return userPrefered;
};

export const setUserPreferedTheme = (newTheme) => {
	if (!newTheme) localStorage.removeItem(localStorageKey);

	localStorage.setItem(localStorageKey, newTheme === "dark" ? "true" : "false");
};
