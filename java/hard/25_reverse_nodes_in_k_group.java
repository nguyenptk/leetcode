// https://leetcode.com/problems/reverse-nodes-in-k-group/

package java.hard;

class ReverseNodesInKGroup {
    public ListNode reverseKGroup(ListNode head, int k) {
        ListNode current = head;
        int count = 0;
        while (count != k && head.next != null) {
            current = current.next;
            count++;
        }
        if (count == k) {
            current = reverseKGroup(current, k);
            while (count > 0) {
                ListNode tmp = head.next;
                head.next = current;
                current = head;
                head = tmp;
                count--;
            }
            head = current;
        }

        return head;
    }
}
