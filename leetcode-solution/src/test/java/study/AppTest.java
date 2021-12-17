package study;

import junit.framework.Test;
import junit.framework.TestCase;
import junit.framework.TestSuite;

/**
 * Unit test for simple App.
 */
public class AppTest 
    extends TestCase
{
    /**
     * Create the test case
     *
     * @param testName name of the test case
     */
    public AppTest( String testName )
    {
        super( testName );
    }

    /**
     * @return the suite of tests being tested
     */
    public static Test suite()
    {
        return new TestSuite( AppTest.class );
    }

    /**
     * Rigourous Test :-)
     */
    public void testApp()
    {
        assertTrue( true );
    }

    public void testReverseInt() {
        ReverseInt r = new ReverseInt();
        assert 321 == r.reverse(123);
        assert -321 == r.reverse(-123);
        assert 21 == r.reverse(120);
        assert 0 == r.reverse(0);
        assert 0 == r.reverse(-2147483648);
        assert 2147483641 == r.reverse(1463847412);
//        assert -2147483648
//        assert 0 == new ReverseInt().reverse(-2147483648);
    }
}
