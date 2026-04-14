import java.util.Scanner;

class Solution {

	static boolean isSafe(int[] a) {
		boolean isInc = true;
		boolean isDec = true;
		int maxDiff = 0;
		for (int i = 1; i < a.length; i++) {
			if (a[i] >= a[i - 1]) {
				isDec = false;
			}
			if (a[i] <= a[i - 1]) {
				isInc = false;
			}
			maxDiff = Math.max(maxDiff, Math.abs(a[i] - a[i - 1]));

		}
		return (isInc || isDec) && maxDiff <= 3;
	}

	public static void main(String[] args) {
		Scanner s = new Scanner(System.in);
		int cnt1 = 0;
		int cnt2 = 0;
		while (s.hasNext()) {
			String line = s.nextLine();
			String[] lineArr = line.split(" ");
			int[] a = new int[lineArr.length];
			for (int i = 0; i < a.length; i++) {
				a[i] = Integer.parseInt(lineArr[i]);
			}
			if (isSafe(a)) {
				cnt1 += 1;
				cnt2 += 1;
				continue;
			}
			for (int i = 0; i < a.length; i++) {
				int[] b = new int[a.length - 1];
				for (int j = 0; j < b.length; j++) {
					if (j < i) {
						b[j] = a[j];
					} else {
						b[j] = a[j + 1];
					}
				}
				if (isSafe(b)) {
					cnt2 += 1;
					break;
				}
			}
		}
		System.out.println(cnt1);
		System.out.println(cnt2);
		s.close();
	}
}
