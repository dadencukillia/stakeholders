type theme = "light" | "dark";
type preferedTheme = theme | undefined;

export function getUserPreferedTheme(): preferedTheme;
export function getBrowserTheme(): theme;
export function getTheme(): theme;
export function setUserPreferedTheme(newTheme: preferedTheme): void;
