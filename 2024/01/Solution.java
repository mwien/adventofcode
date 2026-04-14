import java.util.Scanner;
import java.util.ArrayList;
import java.util.Collections;
import java.util.HashSet;

class Solution {

	static void part1(List<Integer> first, List<Integer> second) {
		Collections.sort(first);
		Collections.sort(second);
		int res = 0;
		for (int i = 0; i < first.size(); i++) {
			res += Math.abs(first.get(i) - second.get(i));
		}
		System.out.println(res);
	}

	static void part2(List<Integer> first, List<Integer> second) {
		HashSet<Integer> h = new HashSet<Integer>();
		for (int x : first) {
			h.add(x);
		}
		int res = 0;
		for (int y : second) {
			if (h.contains(y)) {
				res += y;
			}
		}
		System.out.println(res);
	}

	public static void main(String[] args) {
		Scanner s = new Scanner(System.in);
		ArrayList<Integer> first = new ArrayList<>();
		ArrayList<Integer> second = new ArrayList<>();
		while (s.hasNext()) {
			first.add(s.nextInt());
			second.add(s.nextInt());
		}
		part1(first, second);
		part2(first, second);
		s.close();
	}
}
