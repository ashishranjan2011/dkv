package org.dkv.client;

import com.google.gson.Gson;
import dkv.serverpb.Api;
import org.junit.After;
import org.junit.Before;
import org.junit.Test;

import java.io.InputStream;
import java.io.InputStreamReader;

import static dkv.serverpb.Api.ReadConsistency.LINEARIZABLE;
import static dkv.serverpb.Api.ReadConsistency.SEQUENTIAL;
import static java.lang.String.format;
import static org.junit.Assert.*;

public class ShardedDKVClientTest {

    private static final String KEY_PREFIX = "key_135";
    private static final int NUM_KEYS = 9000;
    private static final Api.ReadConsistency READ_CONSISTENCY = SEQUENTIAL;

    private ShardedDKVClient dkvClient;

    @Before
    public void setup() {
        ShardConfiguration shardConf = loadShardConfig("/local_dkv_config.json");
//        ShardConfiguration shardConf = loadShardConfig("/local_dkv_config_via_envoy.json");
//        ShardConfiguration shardConf = loadShardConfig("/single_local_dkv_config.json");
        dkvClient = new ShardedDKVClient(new KeyHashBasedShardProvider(shardConf));
    }

    @Test
    public void shouldPerformPutAndGet() {
        String[] keys = new String[NUM_KEYS];
        String[] expVals = new String[NUM_KEYS];
        for (int i = 0; i < NUM_KEYS; i++) {
            keys[i] = format("%s%d", KEY_PREFIX, i);
            expVals[i] = format("val_%d", i);
            dkvClient.put(keys[i], expVals[i]);
        }

        for (int i = 0; i < NUM_KEYS; i++) {
            String actVal = dkvClient.get(READ_CONSISTENCY, keys[i]);
            assertEquals(format("Invalid value for key: %s", keys[i]), expVals[i], actVal);
        }

        String[] actVals = dkvClient.multiGet(READ_CONSISTENCY, keys);
        for (int i = 0; i < actVals.length; i++) {
            String actVal = actVals[i];
            String expVal = expVals[i];
            assertEquals(format("Invalid value for key: %s", keys[i]), expVal, actVal);
        }

        try {
            dkvClient.multiGet(LINEARIZABLE, keys);
        } catch (Exception e) {
            assertTrue(e instanceof UnsupportedOperationException);
        }
    }

    @SuppressWarnings("SameParameterValue")
    private ShardConfiguration loadShardConfig(String configPath) {
        InputStream configStream = this.getClass().getResourceAsStream(configPath);
        assertNotNull(configStream);
        return new Gson().fromJson(new InputStreamReader(configStream), ShardConfiguration.class);
    }

    @After
    public void teardown() {
        dkvClient.close();
    }
}