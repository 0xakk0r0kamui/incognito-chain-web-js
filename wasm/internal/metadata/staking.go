package metadata

import (
	"encoding/json"
	"errors"
	metadataCommon "incognito-chain/metadata/common"
)

type StakingMetadata struct {
	MetadataBase
	FunderPaymentAddress         string
	RewardReceiverPaymentAddress string
	StakingAmount                uint64
	AutoReStaking                bool
	CommitteePublicKey           string
	// CommitteePublicKey PublicKeys of a candidate who join consensus, base58CheckEncode
	// CommitteePublicKey string <= encode byte <= mashal struct
}

// NewStakingMetadata creates a new StakingMetadata.
func NewStakingMetadata(
	stakingType int,
	funderPaymentAddress string,
	rewardReceiverPaymentAddress string,
	stakingAmountShard uint64,
	committeePublicKey string,
	autoReStaking bool,
) (
	*StakingMetadata,
	error,
) {
	if stakingType != ShardStakingMeta && stakingType != BeaconStakingMeta {
		return nil, errors.New("invalid staking type")
	}
	metadataBase := NewMetadataBase(stakingType)
	return &StakingMetadata{
		MetadataBase:                 *metadataBase,
		FunderPaymentAddress:         funderPaymentAddress,
		RewardReceiverPaymentAddress: rewardReceiverPaymentAddress,
		StakingAmount:                stakingAmountShard,
		CommitteePublicKey:           committeePublicKey,
		AutoReStaking:                autoReStaking,
	}, nil
}

// GetType overrides MetadataBase.GetType().
func (stakingMetadata StakingMetadata) GetType() int {
	return stakingMetadata.Type
}

// CalculateSize overrides MetadataBase.CalculateSize().
func (stakingMetadata *StakingMetadata) CalculateSize() uint64 {
	return calculateSize(stakingMetadata)
}

func (stakingMetadata StakingMetadata) GetShardStateAmount() uint64 {
	return stakingMetadata.StakingAmount
}

func (sm *StakingMetadata) UnmarshalJSON(raw []byte) error {
	var temp struct {
		MetadataBase
		FunderPaymentAddress         string
		RewardReceiverPaymentAddress string
		StakingAmount                metadataCommon.Uint64Reader
		AutoReStaking                bool
		CommitteePublicKey           string
	}
	err := json.Unmarshal(raw, &temp)
	*sm = StakingMetadata{
		MetadataBase:                 temp.MetadataBase,
		FunderPaymentAddress:         temp.FunderPaymentAddress,
		RewardReceiverPaymentAddress: temp.RewardReceiverPaymentAddress,
		StakingAmount:                uint64(temp.StakingAmount),
		AutoReStaking:                temp.AutoReStaking,
		CommitteePublicKey:           temp.CommitteePublicKey,
	}
	return err
}
